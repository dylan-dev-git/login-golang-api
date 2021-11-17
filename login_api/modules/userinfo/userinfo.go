package userinfo

import (
	"context"
	"errors"
	"net/http"

	"github.com/dylan-dev-git/login-microservice-api/dbconnector"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	errNotExist = errors.New("Not exist email address")
	errTokenVerifyFailed = errors.New("Error token verify")
)

func GetUserInfo(c *gin.Context) {
	userid := c.Request.Header.Get("Userid")
	userinfoDB := dbconnector.GetUserInfoCollection()
	filter := bson.M{"email": userid}
	var result bson.M
	err := userinfoDB.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": errNotExist.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "userData": &result})
}