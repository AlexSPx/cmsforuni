package main

import (
	"fmt"
	"log"

	"github.com/alexspx/gocms/src/initializers"
	"github.com/alexspx/gocms/src/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDatabase(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
