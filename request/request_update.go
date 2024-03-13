package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/naufaldinta13/cars/controllers"
	"github.com/naufaldinta13/cars/entity"
)

type UpdateRequest struct {
	ID        string  `json:"-" validate:"id"`
	CarName   string  `json:"car_name" binding:"required" validate:"unique"`
	DayRate   float64 `json:"day_rate" binding:"required"`
	MonthRate float64 `json:"month_rate" binding:"required"`
	Image     string  `json:"image" binding:"required"`

	Car *entity.Car `json:"-"`
}

func (r *UpdateRequest) Validate() (e error) {
	v := validator.New()

	v.RegisterValidation("id", func(fl validator.FieldLevel) bool {
		if r.Car, e = controllers.NewRentRepository().Show(r.ID); e != nil {
			return false
		}

		return true
	})

	v.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		if r.CarName != "" {
			if _, e = controllers.NewRentRepository().FindByName(r.CarName, r.ID); e == nil {
				return false
			}
		}

		return true
	})

	return v.Struct(r)
}

func (r *UpdateRequest) Execute() (mx *entity.Car, e error) {
	mx = &entity.Car{
		ID:        r.Car.ID,
		CarName:   r.CarName,
		DayRate:   r.DayRate,
		MonthRate: r.MonthRate,
		Image:     r.Image,
		IsDeleted: r.Car.IsDeleted,
	}

	mx, e = controllers.NewRentRepository().Update(mx)

	return
}
