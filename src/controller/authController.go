package controller

import (
	"go-gin-api/src/entity"
	"go-gin-api/src/logger"
	"go-gin-api/src/middleware"
	"go-gin-api/src/model"
	"go-gin-api/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type authController struct {
	service     service.IAuthService
	logger      logger.Logrus
	userService service.IUserService
}

func NewAuth(server *gin.Engine, service service.IAuthService, jwtService service.IJWTService, userService service.IUserService, logger logger.Logrus) {

	controller := &authController{
		service:     service,
		userService: userService,
		logger:      logger,
	}

	server.POST("/login", func(ctx *gin.Context) {

		token, err := controller.login(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"accessToken": token})
		}
	})

	server.POST("/register", func(ctx *gin.Context) {

		message, err := controller.register(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": message})
		}
	})

	server.GET("/me", func(ctx *gin.Context) {
		middleware.AuthMiddleware(ctx, jwtService, logger)
	}, func(ctx *gin.Context) {

		user, err := controller.me(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": user})
		}
	})

}

func (c *authController) register(ctx *gin.Context) (string, error) {

	var userObj entity.RegisterUser
	err := ctx.ShouldBindJSON(&userObj)
	if err != nil {
		return "", err
	}

	_, err = c.service.Register(&userObj)

	if err != nil {
		return "", err
	}

	return "user register successfully", nil
}

func (c *authController) login(ctx *gin.Context) (string, error) {

	var userCred entity.UserCredentials
	err := ctx.ShouldBindJSON(&userCred)
	if err != nil {
		return "", err
	}

	token, err := c.service.Login(&userCred)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (c *authController) me(ctx *gin.Context) (*model.User, error) {

	hexString := ctx.Request.Header.Get("userId")
	userId, err := primitive.ObjectIDFromHex(hexString)
	if err != nil {
		return nil, err
	}

	userObj, err := c.userService.FindOneById(userId)
	if err != nil {
		return nil, err
	}
	return userObj, nil
}
