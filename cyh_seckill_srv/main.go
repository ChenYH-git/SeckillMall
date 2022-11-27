package main

import (
	"cyh_project/cyh_seckill_srv/controller"
	_ "cyh_project/cyh_seckill_srv/data_source"
	"cyh_project/cyh_seckill_srv/redis_lib"
)

var (
	service = "cyh_seckill_srv"
	version = "latest"
	address = ":8083"
)

func main() {
	redis_lib.InitRedis()
	defer redis_lib.CloseRedis()
	controller.InitConsumerMq()
	// Create service
	//registry := consul.NewRegistry()
	//rpcServer := server.NewServer(
	//	server.Name(service),
	//	server.Version(version),
	//	server.Address(address),
	//	server.Registry(registry),
	//)
	//
	//srv := micro.NewService(micro.Server(rpcServer))
	//srv.Init()
	//
	//// Register handler
	//if err := seckill.RegisterSecKillHandler(srv.Server(), new(controller.SecKill)); err != nil {
	//	logger.Fatal(err)
	//}
	//// Run service
	//if err := srv.Run(); err != nil {
	//	logger.Fatal(err)
	//}
}
