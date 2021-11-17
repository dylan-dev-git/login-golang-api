package token

import (
	"errors"
	"net/http"

	"github.com/dylan-dev-git/token-microservice-api/jwt"
	"github.com/gin-gonic/gin"
)

var (
	errTokenIssueFailed  = errors.New("error issue token")
	errTokenVerifyFailed = errors.New("error token verify")
)

func GetToken(c *gin.Context) {
	userID := c.Param("userid")

	token, err := jwt.CreateToken(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": errTokenIssueFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "token": token, "userid": userID, "message": "Token issued."})
}

func TokenVaildChecker(c *gin.Context) {
	_, err := jwt.VerifyToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": errTokenVerifyFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Vaild token."})
}

func GetTokenMetaData(c *gin.Context) {
	userID, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "userid": "", "message": errTokenVerifyFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "userid": userID, "message": "Vaild token."})
}