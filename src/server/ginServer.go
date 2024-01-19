package server

import (
	"go-gin-api/src/config"

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

	loadControllersFn(c)

	// server.Use(gin.Recovery())
	// server.Use(gin.Logger())
	// or
	server.Use(gin.Recovery(), gin.Logger())

	return &ginServer{server: server}
}

func (g *ginServer) Start() {

	port := config.GetConfig().Port
	if port == "" {
		port = "8080"
	}
	g.server.Run(":" + port)
}
