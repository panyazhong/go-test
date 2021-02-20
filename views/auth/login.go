package auth

import (
	"dapan/dbx"
	"dapan/model"
	"dapan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	type Logininfo struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	

	var login_info Logininfo

	err := c.ShouldBind(&login_info)
	if err != nil {
		c.JSON(422, gin.H{
			"err": err.Error(),
		})
	} 

	db := dbx.DB
	var u model.UserInfo

	db.Where("username = ?", login_info.Username).First(&u)

	if (u.ID == 0) {
		c.JSON(http.StatusOK, gin.H{
			"message": "该用户不存在",
		})
		return
	}

	passwordOK := u.CheckPwd(login_info.Password)

	if !passwordOK {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "密码错误",
		})
		return
	}

	token, tokenErr := utils.GenToken(u.ID)

	if tokenErr != nil {
		c.JSON(422, gin.H{
			"err": tokenErr.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}