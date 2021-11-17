package jwt

import (
	"os"
	"strings"
	"time"
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/dylan-dev-git/token-microservice-api/env"
	"github.com/gin-gonic/gin"
)

// Generate Token
func CreateToken(userID string) (string, error) {
	if userID == "" {
		return "", errors.New("UserID is null")
	}
	endOfDay := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, time.Now().Nanosecond(), time.Now().Location()).Unix()
	// twoHours := time.Now().Add(time.Minute * 2).Unix()
	os.Setenv("ACCESS_SECRET", env.JWT_ACCESS_SECRET)
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	// atClaims["exp"] = twoHours
	atClaims["exp"] = endOfDay
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

// token verify and extract data
func ExtractToken(c *gin.Context) string {
	bearToken := c.Request.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenMetadata(c *gin.Context) (string, error) {
	token, err := VerifyToken(c)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, ok := claims["user_id"].(string)
		if !ok {
			return "", err
		}
		return userID, err
	}
	return "", err
}
