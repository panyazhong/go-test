package model

type UserRole struct {
	ID    uint   `form:"id"`
	Name  string `form:"name"`
	Value string `form:"value" gorm:"unique"`
}