package controllers

import (
	"github.com/jinzhu/gorm"

	"github.com/naufaldinta13/cars/config"
	"github.com/naufaldinta13/cars/entity"
)

type RentRepository struct {
	db *gorm.DB
}

func NewRentRepository() *RentRepository {
	return &RentRepository{
		db: config.GetDB(),
	}
}

func (r *RentRepository) Create(m *entity.Car) (mx *entity.Car, e error) {
	e = r.db.Create(m).Error

	return m, e
}

func (r *RentRepository) Update(m *entity.Car, fields ...string) (mx *entity.Car, e error) {
	e = r.db.Model(&m).UpdateColumn(fields).Error

	return m, e
}

func (r *RentRepository) Show(id string) (mx *entity.Car, e error) {
	var m entity.Car
	if e = r.db.Where("id = ?", id).First(&m).Error; e == nil {
		return &m, nil
	}

	return
}

func (r *RentRepository) Delete(m *entity.Car) (e error) {
	m.IsDeleted = 1
	return r.db.Save(m).Error
}

func (r *RentRepository) FindByName(name string, exclude string) (mx *entity.Car, e error) {
	var m entity.Car
	if e = r.db.Raw("SELECT * FROM cars WHERE car_name = ? AND is_deleted = 0 AND id != ?", name, exclude).Scan(&m).Error; e == nil {
		return &m, nil
	}

	return
}
