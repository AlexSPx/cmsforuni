package routes

import (
	"github.com/alexspx/gocms/src/controllers"
	"github.com/alexspx/gocms/src/middlewares"
	"github.com/gin-gonic/gin"
)

type AuthRouterController struct {
	authController controllers.AuthController
}

func NewAuthRouterController(authController controllers.AuthController) AuthRouterController {
	return AuthRouterController{authController}
}

func (rc *AuthRouterController) AuthRoutes(group *gin.RouterGroup) {
	router := group.Group("/auth")

	router.POST("/signin", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	router.GET("/refresh", rc.authController.RefreshAccessToken)
	router.GET("/logout", middlewares.GetUser(), rc.authController.Logout)
}
