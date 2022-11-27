package main

import (
	"cyh_project/cyh_web/middleware"
	all_router "cyh_project/cyh_web/router"

	"go-micro.dev/v4/web"

	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/logger"
)

var (
	service = "cyh_web"
	version = "latest"
	address = ":8081"
)

func main() {
	// Create service
	router := gin.Default()
	router.Use(middleware.CorsMiddleWare)

	all_router.InitRouter(router)
	// 使用全局中间件，跨域请求

	registry := consul.NewRegistry()
	srv := web.NewService(
		web.Name(service),
		web.Version(version),
		web.Handler(router),
		web.Address(address),
		web.Registry(registry),
	)

	if err := srv.Init(); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
