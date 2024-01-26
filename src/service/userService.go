package service

import (
	"go-gin-api/src/dao"
	"go-gin-api/src/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserService interface {
	FindOneById(*gin.Context, primitive.ObjectID) (*model.User, error)
}

type userService struct {
	jwtService IJWTService
	userDao    dao.IUserDao
}

func NewUserService(jwtService IJWTService, ud dao.IUserDao) IUserService {
	return &userService{
		jwtService: jwtService,
		userDao:    ud,
	}
}

func (s *userService) FindOneById(ctx *gin.Context, userId primitive.ObjectID) (*model.User, error) {
	return findUserByObjectId(ctx, s.userDao, userId)
}

// this method can b invoke within service package
func findUserByObjectId(ctx *gin.Context, userDao dao.IUserDao, userId primitive.ObjectID) (*model.User, error) {
	userObj, err := userDao.FindOne(bson.M{"_id": userId})

	if err != nil {
		return nil, err
	}
	return userObj, nil
}

func getLoggedInUser(ctx *gin.Context, userDao dao.IUserDao) (*model.User, error) {
	// get user object Id
	hexString := ctx.Request.Header.Get("userId")
	userId, err := primitive.ObjectIDFromHex(hexString)
	if err != nil {
		return nil, err
	}

	// fetch user data from mongo
	userObj, err := findUserByObjectId(ctx, userDao, userId)
	if err != nil {
		return nil, err
	}

	return userObj, nil
}
