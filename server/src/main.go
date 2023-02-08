package main

import (
	"log"

	"github.com/alexspx/gocms/src/initializers"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
	config *initializers.Config
)

func init() {
	config, err := initializers.LoadConfig("/")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDatabase(&config)

	server = gin.Default()
}

func main() {
	router := server.Group("/api")

	router.Group("/api")

	log.Fatal(server.Run(":" + config.ServerPort))
}
