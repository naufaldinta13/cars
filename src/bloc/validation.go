package bloc

import (
	"github.com/env-io/orm"
	"github.com/naufaldinta13/cars/entity"
)

func ValidDuplicate(name string, exclude string) bool {
	var tot int64
	orm.NewOrm().Raw("SELECT count(id) FROM cars WHERE name = ? AND car_id != ?", name, exclude).QueryRow(&tot)

	return tot == 0
}

func ValidID(id string) (mx *entity.Cars, e error) {
	e = orm.NewOrm().Raw("SELECT * FROM cars where id = ?", id).QueryRow(&mx)

	return
}
