package users

import (
	"github.com/gin-gonic/gin"
)

func UsersRoute(publicApi *gin.RouterGroup) {
	usersGroup := publicApi.Group("/users")
	{
		usersGroup.GET("/list", GetUserList)
	}
}