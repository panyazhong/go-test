package users

import (
	"dapan/dbx"
	"dapan/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	db := dbx.DB

	var u []model.UserInfo
	db.Find(&u)

	c.JSON(http.StatusOK, gin.H{
		"info": u,
	})
}