package router

import (
	"dapan/views/users"

	"github.com/gin-gonic/gin"
)

func SetView(r *gin.Engine) {
	publicApi := r.Group("/api")

	users.UsersRoute(publicApi)
}