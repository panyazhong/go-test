package auth

import (
	// "dapan/dbx"
	"dapan/dbx"
	"dapan/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var u model.UserInfo

	err := c.ShouldBind(&u)

	if err != nil {
		c.JSON(422, gin.H{
			"message": "参数错误",
		})
	}

	u.SetPassword(u.Password)

	dbx.DB.Create(&u)

	c.JSON(200, gin.H{
		"message": "新建成功",
	})
}