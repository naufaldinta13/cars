package rest

import (
	"github.com/naufaldinta13/cars/entity"
	"github.com/naufaldinta13/cars/src/bloc"

	"github.com/env-io/orm"
	"github.com/env-io/validate"
)

type updateRequest struct {
	ID        string  `json:"-"`
	Name      string  `json:"name" valid:"required"`
	MonthRate float64 `json:"month_rate" valid:"required"`
	Image     string  `json:"image" valid:"required"`

	Car *entity.Cars `json:"-"`
}

func (r *updateRequest) Validate() *validate.Response {
	v := validate.NewResponse()

	var e error

	if bloc.ValidDuplicate(r.Name, r.ID) {
		v.SetError("name.required", "nama sudah tersedia.")
	}

	if r.Car, e = bloc.ValidID(r.ID); e != nil {
		v.SetError("id.invalid", "data tidak ditemukan.")
	}

	return v
}

func (r *updateRequest) Messages() map[string]string {
	return map[string]string{
		"name.required": "nama harus diisi.",
	}
}

func (r *updateRequest) Execute() (m *entity.Cars, e error) {
	m = &entity.Cars{
		CarID:     r.Car.CarID,
		CarName:   r.Name,
		MonthRate: r.MonthRate,
		Image:     r.Image,
	}

	if _, e = orm.NewOrm().Update(m, "car_name", "month_rate", "image"); e != nil {
		return
	}

	return
}
