package protos

import "github.com/naufaldinta13/cars/entity"

func ConvertCarProto(m *entity.Car) *Car {
	return &Car{
		Id:        m.ID,
		CarName:   m.CarName,
		DayRate:   m.DayRate,
		MonthRate: m.MonthRate,
		Image:     m.Image,
	}
}

func ConvertCarEntity(m *Car) *entity.Car {
	return &entity.Car{
		ID:        m.Id,
		CarName:   m.CarName,
		DayRate:   m.DayRate,
		MonthRate: m.MonthRate,
		Image:     m.Image,
	}
}
