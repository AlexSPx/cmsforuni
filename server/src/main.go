package main

import (
	"log"
	"net/http"

	"github.com/alexspx/gocms/src/controllers"
	"github.com/alexspx/gocms/src/initializers"
	"github.com/alexspx/gocms/src/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
	config initializers.Config

	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouterController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
)

func init() {
	var err error
	config, err = initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDatabase(&config)

	server = gin.Default()

	AuthController = controllers.NewAuthController(initializers.DB, &config)
	AuthRouteController = routes.NewAuthRouterController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)
}

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "+opened api"})
	})

	AuthRouteController.AuthRoutes(router)
	UserRouteController.UserRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
