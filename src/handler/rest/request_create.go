package rest

import (
	"github.com/naufaldinta13/cars/entity"
	"github.com/naufaldinta13/cars/src/bloc"

	"github.com/env-io/orm"
	"github.com/env-io/validate"
)

type createRequest struct {
	Name      string  `json:"name" valid:"required"`
	MonthRate float64 `json:"month_rate"`
	Image     string  `json:"image"`
}

func (r *createRequest) Validate() *validate.Response {
	v := validate.NewResponse()

	if bloc.ValidDuplicate(r.Name, "") {
		v.SetError("name.required", "nama sudah tersedia.")
	}

	return v
}

func (r *createRequest) Messages() map[string]string {
	return map[string]string{
		"name.required": "nama harus diisi.",
	}
}

func (r *createRequest) Execute() (m *entity.Cars, e error) {
	m = &entity.Cars{
		CarName:   r.Name,
		MonthRate: r.MonthRate,
		Image:     r.Image,
		IsDeleted: false,
	}

	if m.CarID, e = orm.NewOrm().Insert(m); e != nil {
		return
	}

	return
}
