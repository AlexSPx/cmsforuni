package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(config *Config) {
	var error error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	DB, error = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if error != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("? Connected Successfully to the Database")
}
