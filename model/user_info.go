package model

import (
	"golang.org/x/crypto/bcrypt"
)

type UserInfo struct {
	ID 			 uint 	`form:"id"`
	Username string `form:"username" binding:"required" gorm:"type:varchar(128)"`
	Password string `form:"password" binding:"required" gorm:"type:varchar(128)"`
	Gender 	 string `form:"gender"`
	RoleId 	 uint 	`form:"role_id"`
}

func (u *UserInfo) SetPassword(password string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if (err != nil) {
		return
	}

	u.Password = string(hash)

	return
}


func (u *UserInfo) CheckPwd(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return nil == err
}