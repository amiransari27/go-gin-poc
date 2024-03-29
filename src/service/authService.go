package service

import (
	"fmt"
	"go-gin-api/src/dao"
	"go-gin-api/src/entity"
	"go-gin-api/src/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(*gin.Context, *entity.UserCredentials) (string, error)
	Register(*gin.Context, *entity.RegisterUser) (interface{}, error)
}

type authService struct {
	jwtService IJWTService
	userDao    dao.IUserDao
}

func NewAuth(jwtService IJWTService, ud dao.IUserDao) IAuthService {
	return &authService{
		jwtService: jwtService,
		userDao:    ud,
	}
}

func (service *authService) Login(ctx *gin.Context, userCred *entity.UserCredentials) (string, error) {

	// check user exist or not
	userObj, err := service.userDao.FindOne(bson.M{"username": userCred.Username})

	if err != nil {
		return "", err
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(userObj.Password), []byte(userCred.Password))

	if err != nil {
		return "", err
	}
	// create jwt token
	token, err := service.jwtService.GenerateJWTToken(userObj.Id.Hex())

	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *authService) Register(ctx *gin.Context, userObj *entity.RegisterUser) (interface{}, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userObj.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userId, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	newuser := &model.User{
		UserId:    userId.String(),
		Username:  userObj.Username,
		Password:  string(hashedPassword),
		FirstName: userObj.FirstName,
		LastName:  userObj.LastName,
	}

	inserted, err := service.userDao.Save(newuser)

	if err != nil {
		return nil, err
	}

	fmt.Println("user ", inserted)

	return inserted, nil
}
