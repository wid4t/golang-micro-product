package main

import (
	"golang-micro-product/handler"
	pb "golang-micro-product/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	serviceName = "golang-micro-product"
	version     = "latest"
)

func main() {
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version(version),
	)
	service.Init()

	if err := pb.RegisterGolangMicroProductHandler(service.Server(), new(handler.GolangMicroProduct)); err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
