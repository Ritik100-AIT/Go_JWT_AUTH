package routes

import (
	"github.com/Ritik100-AIT/Go_JWT_AUTH/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) {
	route.POST("users/signup", controllers.Signup())
	route.POST("users/login", controllers.Login())
}
