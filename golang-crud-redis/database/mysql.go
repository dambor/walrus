package database

import (
	"fmt"
	"go-crud-redis-example/config"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func ConnectionMySQLDb(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	fmt.Println("Connected Successfuly to the database (MySQL)")

	return db
}
