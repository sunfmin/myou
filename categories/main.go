package main

import (
	"log"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/transport/nats"

	"github.com/sunfmin/myou/categories/handler"

	cats "github.com/sunfmin/myou/categories/proto/categories"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("myou.srv.categories"),
		micro.Version("latest"),
	)

	// Register Handler
	cats.RegisterCategoriesHandler(service.Server(), new(handler.CatHandler))

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
