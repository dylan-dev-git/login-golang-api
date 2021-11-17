package routes

import (
	"net/http"

	"github.com/dylan-dev-git/token-microservice-api/env"
	"github.com/gin-gonic/gin"

	tokenMod "github.com/dylan-dev-git/token-microservice-api/modules/token"
)

// middleware for API KEY and VALUE check.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get(env.API_KEY) != env.API_VALUE {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Token API authorication failed"})
			return
		}
		c.Next()
	}
}

func StartGin() {
	router := gin.Default()
	api := router.Group("/api")
	api.Use(AuthRequired())
	{
		token := api.Group("/token")
		{
			token.GET("/getToken/:userid", tokenMod.GetToken)
			token.GET("/tokenVaildChecker", tokenMod.TokenVaildChecker)
			token.GET("/getTokenMetaData", tokenMod.GetTokenMetaData)
		}
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8501")
}
