package rest

import (
	"github.com/naufaldinta13/cars/entity"
	"github.com/naufaldinta13/cars/src/bloc"

	"github.com/env-io/orm"
	"github.com/env-io/validate"
)

type deleteRequest struct {
	ID string `json:"-"`

	Car *entity.Cars `json:"-"`
}

func (r *deleteRequest) Validate() *validate.Response {
	v := validate.NewResponse()

	var e error

	if r.Car, e = bloc.ValidID(r.ID); e != nil {
		v.SetError("id.invalid", "data tidak ditemukan.")
	}

	return v
}

func (r *deleteRequest) Messages() map[string]string {
	return map[string]string{}
}

func (r *deleteRequest) Execute() (e error) {
	_, e = orm.NewOrm().Raw("UPDATE cars SET is_deleted = true WHERE id = ?", r.ID).Exec()

	return
}
