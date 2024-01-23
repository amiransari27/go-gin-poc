package service

import (
	"fmt"
	"go-gin-api/src/logger"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var tokenHour = "24"
var apiSecret = "yoursecretstring"

type IJWTService interface {
	GenerateJWTToken(string) (string, error)
	TokenValid(*gin.Context) error
}

type jwtService struct {
	logger logger.ILogrus
}

func NewJWT(logger logger.ILogrus) IJWTService {
	return &jwtService{
		logger,
	}
}

func (service *jwtService) GenerateJWTToken(userId string) (string, error) {

	tokenLifespan, err := strconv.Atoi(tokenHour)

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(apiSecret))

}

func (service *jwtService) TokenValid(ctx *gin.Context) error {
	tokenString := extractToken(ctx)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(apiSecret), nil
	})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId := fmt.Sprintf("%s", claims["userId"])
		ctx.Request.Header.Set("userId", userId)
		return nil
	}

	return nil
}
func extractToken(ctx *gin.Context) string {
	token := ctx.Query("access-token")
	if token != "" {
		return token
	}
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
