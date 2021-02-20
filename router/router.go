package router

import (
	"dapan/utils"
	"dapan/views/auth"
	"dapan/views/users"

	"github.com/gin-gonic/gin"
)

func SetView(r *gin.Engine) {
	publicApi := r.Group("/api")
	auth.AuthRoute(publicApi)

	// 需要认证的路由
	authApi := r.Group("/api")
	authApi.Use(utils.Auth)

	users.UsersRoute(authApi)
}