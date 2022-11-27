package main

import (
	"cyh_project/cyh_user_srv/controller"
	_ "cyh_project/cyh_user_srv/data_source"
	cyh_admin_srv "cyh_project/cyh_user_srv/proto/admin_user"
	cyh_user_srv "cyh_project/cyh_user_srv/proto/front_user"

	"go-micro.dev/v4/server"

	"github.com/go-micro/plugins/v4/registry/consul"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "cyh_user_srv"
	version = "latest"
	address = ":8082"
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
	if err := cyh_user_srv.RegisterFrontUserHandler(srv.Server(), new(controller.FrontUser)); err != nil {
		logger.Fatal(err)
	}
	if err := cyh_admin_srv.RegisterAdminUserHandler(srv.Server(), new(controller.AdminUser)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
