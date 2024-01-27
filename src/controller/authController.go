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
	logger      logger.ILogrus
	userService service.IUserService
}

func NewAuth(server *gin.Engine, service service.IAuthService, jwtService service.IJWTService, userService service.IUserService, logger logger.ILogrus) {

	controller := &authController{
		service:     service,
		userService: userService,
		logger:      logger,
	}

	server.POST("/login", func(ctx *gin.Context) {

		tokenObj, err := controller.login(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, tokenObj)
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

	server.GET("/me", middleware.AuthMiddleware(jwtService, logger), func(ctx *gin.Context) {

		user, err := controller.me(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	})

}

// @Summary Register user
// @Schemes
// @Description register new user
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {string} Register
// @Router /register [post]
// @Param data body entity.RegisterUser true "Register payload"
func (c *authController) register(ctx *gin.Context) (string, error) {

	var userObj entity.RegisterUser
	err := ctx.ShouldBindJSON(&userObj)
	if err != nil {
		return "", err
	}

	_, err = c.service.Register(ctx, &userObj)

	if err != nil {
		return "", err
	}

	return "user register successfully", nil
}

// @Summary Login user
// @Schemes
// @Description login user
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} entity.LoginResponse
// @Router /login [post]
// @Param data body entity.UserCredentials true "Login payload"
func (c *authController) login(ctx *gin.Context) (*entity.LoginResponse, error) {

	var userCred entity.UserCredentials
	err := ctx.ShouldBindJSON(&userCred)
	if err != nil {
		return nil, err
	}

	token, err := c.service.Login(ctx, &userCred)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{AccessToken: token}, nil
}

// @Summary Me
// @Schemes
// @Description fetch logged in user data
// @Tags Auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.User
// @Router /me [get]
// @Param Authorization header string true "Bearer Token"
func (c *authController) me(ctx *gin.Context) (*model.User, error) {

	hexString := ctx.Request.Header.Get("userId")
	userId, err := primitive.ObjectIDFromHex(hexString)
	if err != nil {
		return nil, err
	}

	userObj, err := c.userService.FindOneById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return userObj, nil
}
