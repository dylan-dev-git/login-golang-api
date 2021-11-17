package routes

import (
	"net/http"

	authMod "github.com/dylan-dev-git/login-microservice-api/modules/auth"
	userinfoMod "github.com/dylan-dev-git/login-microservice-api/modules/userinfo"
	token "github.com/dylan-dev-git/login-microservice-api/microservices/token"

	"github.com/dylan-dev-git/login-microservice-api/env"
	"github.com/gin-gonic/gin"
)

// middleware API KEY and VALUE check.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get(env.API_KEY) != env.API_VALUE {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Login API authorication failed."})
			return
		}
		c.Next()
	}
}

// middleware token verify and get metadata.
func TokenMetaData() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenObj := token.NewToken()
		userid, err := tokenObj.GetTokenMetaData(c.Request.Header.Get("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
			return
		}
		c.Request.Header.Add("Userid",userid)
		c.Next()
	}
}

// middleware CORS.
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, LOGINAPIKEYHERE")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func StartGin() {
	router := gin.Default()
	router.Use(Cors())
	api := router.Group("/api")
	api.Use(AuthRequired())
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authMod.Login)
		}
		user := api.Group("/user")
		user.Use(TokenMetaData())
		{
			user.GET("/userinfo", userinfoMod.GetUserInfo)
		}
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":7501")
}
