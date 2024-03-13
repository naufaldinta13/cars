package main

import (
	"log"
	"os"

	"github.com/naufaldinta13/cars/config"
	"github.com/naufaldinta13/cars/grpc"
	"github.com/naufaldinta13/cars/protos"
	"github.com/naufaldinta13/cars/routes"
	"github.com/oklog/run"

	"github.com/asim/go-micro/v3"
	"github.com/joho/godotenv"
)

var Routine run.Group

func initMySQLConnection() {
	c := &config.DBCOnfig{
		Server:   os.Getenv("MYSQL_SERVER"),
		Username: os.Getenv("MYSQL_USERNAME"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Database: os.Getenv("MYSQL_DATABASE"),
	}

	if e := config.NewDBConnection(c); e != nil {
		log.Fatal(e)
	}
}

func initGrpcConnection() {
	c := &config.GrpcConfig{
		Name:           "testing.car",
		RegistryServer: os.Getenv("SERVICE_REGISTRY"),
		Server:         os.Getenv("SERVICE_SERVER"),
	}

	if e := config.NewGrpcConnection(c, func(s micro.Service) {
		protos.RegisterCarServiceHandler(s.Server(), new(grpc.CarService))
	}); e != nil {
		log.Fatal(e)
	}
}

func init() {
	godotenv.Load()

	initMySQLConnection()
	initGrpcConnection()
}

func main() {
	route := routes.SetupRoutes()

	go func() {
		route.Run(os.Getenv("REST_SERVER"))
	}()

	Routine.Add(config.Start, config.Shutdown)
	Routine.Run()
}
