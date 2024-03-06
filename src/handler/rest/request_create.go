package rest

import (
	"github.com/naufaldinta13/cars/entity"
	"github.com/naufaldinta13/cars/src/bloc"

	"github.com/env-io/orm"
	"github.com/env-io/validate"
)

type createRequest struct {
	Name      string  `json:"name" valid:"required"`
	DayRate   float64 `json:"day_rate" valid:"required|gt:0"`
	MonthRate float64 `json:"month_rate" valid:"required|gt:0"`
	Image     string  `json:"image" valid:"required"`
}

func (r *createRequest) Validate() *validate.Response {
	v := validate.NewResponse()

	if !bloc.ValidDuplicate(r.Name, "") {
		v.SetError("name.invalid", "nama sudah tersedia.")
	}

	return v
}

func (r *createRequest) Messages() map[string]string {
	return map[string]string{}
}

func (r *createRequest) Execute() (m *entity.Cars, e error) {
	m = &entity.Cars{
		CarName:   r.Name,
		DayRate:   r.DayRate,
		MonthRate: r.MonthRate,
		Image:     r.Image,
	}

	if m.ID, e = orm.NewOrm().Insert(m); e != nil {
		return
	}

	return
}
