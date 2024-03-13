package grpc

import (
	"context"

	"github.com/naufaldinta13/cars/controllers"
	"github.com/naufaldinta13/cars/protos"
)

type CarService struct{}

func (s *CarService) Show(ctx context.Context, req *protos.ShowRequest, res *protos.CarResponse) error {
	data, e := controllers.NewRentRepository().Show(req.Id)

	if e != nil || data == nil {
		return e
	}

	res.Car = protos.ConvertCarProto(data)

	return e
}
