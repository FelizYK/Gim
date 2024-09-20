package logic

import (
	"Gim/internal/sql"

	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	Username  string
	Password  string
	Salt      string
	Telephone string
	Email     string
	Token     string
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

func GetUserByName(username string) (UserInfo, error) {
	var user UserInfo
	err := sql.DB.Where("username = ?", username).First(&user).Error
	return user, err
}
func GetUserByTel(telephone string) (UserInfo, error) {
	var user UserInfo
	err := sql.DB.Where("telephone = ?", telephone).First(&user).Error
	return user, err
}
func GetUserByEmail(email string) (UserInfo, error) {
	var user UserInfo
	err := sql.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func GetUserList() ([]*UserInfo, error) {
	var users []*UserInfo
	err := sql.DB.Find(&users).Error
	return users, err
}

func UpdateUser(user UserInfo) error {
	return sql.DB.Save(&user).Error
}

func DeleteUser(user UserInfo) error {
	return sql.DB.Delete(&user).Error
}
