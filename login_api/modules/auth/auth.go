package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/dylan-dev-git/login-microservice-api/dbconnector"
	token "github.com/dylan-dev-git/login-microservice-api/microservices/token"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"gopkg.in/hlandau/passlib.v1"
)

var (
	errNotExist          = errors.New("Not exist email address")
	errInvalidBody       = errors.New("Invalid request body")
	errInvalidPassword   = errors.New("Invalid email/password")
	errDisableLogin      = errors.New("Sorry, your ID has expired.")
	errTokenIssueFailed  = errors.New("Token issue error")
)

func Login(c *gin.Context) {
	usersDB := dbconnector.GetUserCollection()

	bodyJSON := make(map[string]interface{})
	err := c.ShouldBind(&bodyJSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": errInvalidBody.Error()})
		return
	}
	loginForm := bodyJSON["loginForm"]
	email := loginForm.(map[string]interface{})["email"].(string)
	password := loginForm.(map[string]interface{})["password"].(string)

	filter := bson.M{"email": email}
	var result bson.M
	err = usersDB.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": errNotExist.Error()})
		return
	}
	if result["expired"] == true {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": errDisableLogin.Error()})
		return
	}
	currentPassword := result["password"].(string)
	newHash, err := passlib.Verify(password, currentPassword)
	if err != nil {
		_ = newHash
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": errInvalidPassword.Error()})
		return
	}

	userid := result["email"].(string)
	tokenObj := token.NewToken()
	tokenValue, err := tokenObj.GetToken(userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "token": tokenValue["token"], "userid": tokenValue["userid"], "message": "Success"})
}