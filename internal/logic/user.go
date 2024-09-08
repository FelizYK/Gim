package logic

import (
	"Gim/internal/sql"

	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	Username  string
	Password  string
	Telephone string
	Email     string
}

func (table *UserInfo) TableName() string {
	return "user_info"
}

func InitUserTable() {
	sql.DB.AutoMigrate(&UserInfo{})
}

// ===== CRUD =====

func CreateUser(user UserInfo) error {
	return sql.DB.Create(&user).Error
}

func GetUser(username, password string) (UserInfo, error) {
	var user UserInfo
	err := sql.DB.Where("username = ? AND password = ?", username, password).First(&user).Error
	return user, err
}

func GetUserList() []*UserInfo {
	var users []*UserInfo
	sql.DB.Find(&users)
	return users
}

func UpdateUser(user UserInfo) error {
	return sql.DB.Save(&user).Error
}

func DeleteUser(user UserInfo) error {
	return sql.DB.Delete(&user).Error
}
