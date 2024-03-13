package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/naufaldinta13/cars/controllers"
	"github.com/naufaldinta13/cars/entity"
)

type CreateRequest struct {
	CarName   string  `json:"car_name" binding:"required" validate:"unique"`
	DayRate   float64 `json:"day_rate" binding:"required"`
	MonthRate float64 `json:"month_rate" binding:"required"`
	Image     string  `json:"image" binding:"required"`
}

func (r *CreateRequest) Validate() (e error) {
	v := validator.New()

	v.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		if r.CarName != "" {
			if _, e = controllers.NewRentRepository().FindByName(r.CarName, ""); e == nil {
				return false
			}
		}

		return true
	})

	return v.Struct(r)
}

func (r *CreateRequest) Execute() (mx *entity.Car, e error) {
	mx = &entity.Car{
		CarName:   r.CarName,
		DayRate:   r.DayRate,
		MonthRate: r.MonthRate,
		Image:     r.Image,
	}

	mx, e = controllers.NewRentRepository().Create(mx)

	return
}
