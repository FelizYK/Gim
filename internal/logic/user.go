package logic

import (
	"Gim/internal/sql"

	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	Id        uint64
	Username  string
	Password  string
	Telephone string
	Email     string
}

func (table *UserInfo) TableName() string {
	return "user_info"
}

func GetUserList() []*UserInfo {
	var users []*UserInfo
	sql.DB.Find(&users)
	return users
}

func CreateUser(user UserInfo) {
	sql.DB.Create(&user)
}

func DeleteUser(user UserInfo) {
	sql.DB.Create(&user)
}
