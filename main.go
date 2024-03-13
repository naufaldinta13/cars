package main

import (
	"fmt"
	"log"
	"os"

	"github.com/naufaldinta13/cars/config"
	"github.com/naufaldinta13/cars/grpc"
	"github.com/naufaldinta13/cars/protos"
	"github.com/naufaldinta13/cars/routes"

	"github.com/asim/go-micro/v3"
	"github.com/joho/godotenv"
)

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
	service := micro.NewService(
		micro.Name("cars"),
		micro.Address(os.Getenv("SERVICE_REGISTRY")),
	)

	service.Init()

	protos.RegisterCarServiceHandler(service.Server(), new(grpc.CarService))

	if err := service.Run(); err != nil {
		fmt.Println(fmt.Sprintf("Failed to GRPC Server %s", err))

		return
	}

	fmt.Println("Connected to GRPC Server")
}

func init() {
	godotenv.Load()

	initMySQLConnection()
	initGrpcConnection()
}

func main() {
	route := routes.SetupRoutes()

	route.Run(os.Getenv("REST_SERVER"))
}
