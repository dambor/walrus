package main

import (
	"fmt"
	"go-crud-redis-example/config"
	"go-crud-redis-example/database"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello, world!")

	// this part we connect to database first
	// lets make env file

	// lets connect to mysql and redis
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Error while loading config", err)

	}
	// mysql
	db := database.ConnectionMySQLDb(&loadConfig)

	// redis
	rdb := database.ConnectionRedisDb(&loadConfig)
	startServer(db, rdb)
}

// startServer
func startServer(db *gorm.DB, rdb *redis.Client) {
	app := fiber.New()

	err := app.Listen(":3400")
	if err != nil {
		panic(err)
	}

}
