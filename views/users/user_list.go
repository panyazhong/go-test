package users

import (
	"dapan/dbx"
	"dapan/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	db := dbx.DB

	// type results
	type User struct {
		Username string
		Name string
		
	}

	var res []User

	var u []model.UserInfo
	db.Table("user_infos").Select("user_infos.username, user_roles.name").Joins("left join user_roles on user_roles.id = user_infos.role_id").Scan(&res)
	// db.Table("user_info").Select("user_info.id as id, user_info.role_id as role_id").Joins("left join user_role on user_role.id = user_info.role_id").Find(&u)
	log.Print(u)
	log.Print(res)
	c.JSON(http.StatusOK, gin.H{
		"data": u,
	})
}