package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
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
		fmt.Println("MySQL connection failed: ", err)
		return
	}
	fmt.Println("MySQL connected ...")
}

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := RDB.Ping(RDB.Context()).Result()
	if err != nil {
		fmt.Println("Redis connection failed: ", err)
		return
	}
	fmt.Println("Redis connected ...")
}

// Publish message to channel
func Publish(ctx context.Context, channel string, message string) error {
	err := RDB.Publish(ctx, channel, message).Err()
	return err
}

// Subscribe to channel
func Subscribe(ctx context.Context, channel string, messages chan<- string) {
	pubSub := RDB.Subscribe(ctx, channel)
	defer pubSub.Close()
	if err := pubSub.Subscribe(ctx, channel); err != nil {
		fmt.Println("Subscribe failed: ", err)
		return
	}

	ch := pubSub.Channel()
	for {
		select {
		case msg := <-ch:
			messages <- msg.Payload
		case <-ctx.Done():
			fmt.Println("Subscription context cancelled or timed out")
			return
		}
	}
}
