package server

import (
	"go-gin-api/docs"
	"go-gin-api/src/config"
	"go-gin-api/src/middleware"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type GinServer interface {
	Start()
}

type ginServer struct {
	server *gin.Engine
}

func NewGinServer(c *dig.Container, loadControllersFn func(c *dig.Container)) GinServer {
	// registering server instance to dig instance
	// since this server instance is getting used in all the controllers to define routes
	c.Provide(gin.New)
	var server *gin.Engine
	err := c.Invoke(func(s *gin.Engine) {
		server = s
	})

	if err != nil {
		panic(err)
	}

	// server.Use(gin.Recovery())
	// server.Use(gin.Logger())
	// or
	server.Use(gin.Recovery(), gin.Logger())
	server.Use(middleware.ClsMiddleware())

	// load controller here.
	loadControllersFn(c)

	docs.SwaggerInfo.BasePath = "/"
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return &ginServer{server: server}
}

func (g *ginServer) Start() {

	port := config.GetConfig().Port
	if port == "" {
		port = "8080"
	}
	g.server.Run(":" + port)
}
