package bloc

import "github.com/env-io/orm"

func ValidDuplicate(name string, exclude string) bool {
	var tot int64
	orm.NewOrm().Raw("SELECT count(id) FROM cars WHERE name = ? AND car_id != ?", name, exclude).QueryRow(&tot)

	return tot == 0
}
