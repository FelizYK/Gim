package server

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
	RD *redis.Client
)

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
	dsn := "root:password@tcp(localhost:3306)/gim_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: mylogger})
	if err != nil {
		panic(err)
	}
	fmt.Println("MySQL connected ...")
}

func InitRedis() {
	RD = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 5,
	})
	_, err := RD.Ping(RD.Context()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis connected ...")
}
