package routes

import (
	"github.com/alexspx/gocms/src/controllers"
	"github.com/alexspx/gocms/src/middlewares"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(group *gin.RouterGroup) {

	router := group.Group("users")
	router.GET("/me", middlewares.GetUser(), uc.userController.GetMe)
}
