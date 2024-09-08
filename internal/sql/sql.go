package sql

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitMySQL() {
	mylogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	var err error
	// DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{})
	dsn := "root:password@tcp(localhost:3306)/gim_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: mylogger})
	if err != nil {
		panic(err)
	}

	fmt.Println("MySQL connected")
}
