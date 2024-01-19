package ioc

import (
	"fmt"
	"go-gin-api/src/controller"
	"go-gin-api/src/dao"
	"go-gin-api/src/logger"
	"go-gin-api/src/mongoDB"

	"go-gin-api/src/service"

	"go.uber.org/dig"
)

var Controllers = []interface{}{
	controller.NewAuth,
	controller.NewStatus,
	controller.NewVideo,
}

// services. repositry, logger, utils
var otherInjectable = []interface{}{
	service.NewAuth,
	service.NewJWT,
	service.NewVideo,
	service.NewUserService,
	logger.NewLogrus,
	mongoDB.NewMongoConnection,

	dao.NewUserDao,
}

func NewKernal() *dig.Container {
	fmt.Print("called Kernal")
	c := dig.New()

	// bind
	for _, injectable := range otherInjectable {
		c.Provide(injectable)
	}

	return c
}

func LoadControllers(c *dig.Container) {
	for _, controller := range Controllers {
		c.Invoke(controller)
	}
}
