package service

import (
	"go-gin-api/src/dao"
	"go-gin-api/src/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserService interface {
	FindOneById(primitive.ObjectID) (*model.User, error)
}

type userService struct {
	jwtService IJWTService
	userDao    dao.UserDao
}

func NewUserService(jwtService IJWTService, ud dao.UserDao) IUserService {
	return &userService{
		jwtService: jwtService,
		userDao:    ud,
	}
}

func (s *userService) FindOneById(userId primitive.ObjectID) (*model.User, error) {
	userObj, err := s.userDao.FindOne(bson.M{"_id": userId})

	if err != nil {
		return nil, err
	}
	return userObj, nil
}
