package protos

import (
	"github.com/naufaldinta13/cars/entity"
)

func ConvertCarResponse(r *CarResponse) (result *entity.Cars, e error) {
	result, e = ConvertCarToEntity(r.Car)

	return
}

func ConvertCarToEntity(m *Car) (result *entity.Cars, e error) {
	result = &entity.Cars{
		ID:        m.Id,
		CarName:   m.CarName,
		MonthRate: m.MonthRate,
		Image:     m.Image,
	}

	return
}

func ConvertCarToProto(m *entity.Cars) *Car {
	return &Car{
		Id:        m.ID,
		CarName:   m.CarName,
		MonthRate: m.MonthRate,
		Image:     m.Image,
	}
}
