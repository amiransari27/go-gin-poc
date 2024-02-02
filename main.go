package main

import (
	"fmt"
	"go-gin-api/docs"
	"go-gin-api/src/config"
	"go-gin-api/src/ioc"
	"go-gin-api/src/middleware"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	goServer "github.com/openscriptsin/go-server"
)

func main() {
	fmt.Println("go gin api")

	kernal := ioc.NewKernal()
	ginServer := goServer.New(kernal, ioc.RegisterControllers)
	ginApp := ginServer.GetEngine()

	ginApp.Use(middleware.GlobalIpBasedRateLimiterMiddleware())
	ginApp.Use(middleware.ClsMiddleware())

	// registering swagger
	docs.SwaggerInfo.BasePath = "/"
	ginApp.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	ginServer.Start(config.GetConfig().Port)

}
