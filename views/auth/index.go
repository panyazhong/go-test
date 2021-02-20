package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRoute(publicApi *gin.RouterGroup) {
	authGroup := publicApi.Group("/auth")
	{
		authGroup.POST("/register", Register)
		authGroup.POST("/login", Login)
	}
}