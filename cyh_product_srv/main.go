package main

import (
	"cyh_project/cyh_product_srv/controller"
	_ "cyh_project/cyh_product_srv/data_source"

	product "cyh_project/cyh_product_srv/proto/product"
	seckill "cyh_project/cyh_product_srv/proto/seckill"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/server"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "cyh_product_srv"
	version = "latest"
	address = ":8084"
)

func main() {
	// Create service
	registry := consul.NewRegistry()
	rpcServer := server.NewServer(
		server.Name(service),
		server.Version(version),
		server.Address(address),
		server.Registry(registry),
	)

	srv := micro.NewService(micro.Server(rpcServer))
	srv.Init()

	// Register handler
	if err := product.RegisterProductsHandler(srv.Server(), new(controller.Products)); err != nil {
		logger.Fatal(err)
	}
	if err := seckill.RegisterSecKillsHandler(srv.Server(), new(controller.SecKills)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
