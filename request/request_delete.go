package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/naufaldinta13/cars/controllers"
	"github.com/naufaldinta13/cars/entity"
)

type DeleteRequest struct {
	ID string `json:"-" validate:"id"`

	Car *entity.Car `json:"-"`
}

func (r *DeleteRequest) Validate() (e error) {
	v := validator.New()

	v.RegisterValidation("id", func(fl validator.FieldLevel) bool {
		if r.Car, e = controllers.NewRentRepository().Show(r.ID); e != nil {
			return false
		}

		if r.Car != nil {
			if r.Car.IsDeleted == 1 {
				return false
			}
		}

		return true
	})

	return v.Struct(r)

}

func (r *DeleteRequest) Execute() (e error) {
	return controllers.NewRentRepository().Delete(r.Car)
}
