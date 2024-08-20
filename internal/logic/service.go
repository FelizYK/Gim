package user

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type User struct {
	ID       int
	Username string
	Password string
}

var DB *gorm.DB

func InitMysql() {
	dsn := "root:password@tcp(localhost:3306)/gim_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

// CRUD

func Create(username, password string) {
	DB.Create(&User{Username: username, Password: password})
}

func Update(id int, username, password string) {
	DB.Model(&User{}).Where("id = ?", id).Updates(User{Username: username, Password: password})
}
