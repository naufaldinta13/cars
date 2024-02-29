package src

import (
	"github.com/naufaldinta13/cars/protos"
	"github.com/naufaldinta13/cars/src/handler/grpc"
	"github.com/naufaldinta13/cars/src/handler/rest"

	"github.com/asim/go-micro/v3"
	"github.com/labstack/echo/v4"
)

func RegisterRestHandler(e *echo.Echo) {
	rest.RegisterHandler(e)
}

func RegisterGrpcHandler(s micro.Service) {
	protos.RegisterCarServiceHandler(s.Server(), new(grpc.CarService))
}
