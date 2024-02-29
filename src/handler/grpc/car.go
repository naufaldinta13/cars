package grpc

import (
	"context"

	"github.com/env-io/orm"
	"github.com/naufaldinta13/cars/entity"
	"github.com/naufaldinta13/cars/protos"
)

type CarService struct{}

func (g *CarService) Show(ctx context.Context, req *protos.ShowRequest, resp *protos.CarResponse) error {
	var mx entity.Cars
	if err := orm.NewOrm().QueryTable("cars").Filter("id", req.Id).One(&mx); err != nil {
		return err
	}

	resp.Car = protos.ConvertCarToProto(&mx)

	return nil
}
