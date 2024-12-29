package routes

import (
	"github.com/Ritik100-AIT/Go_JWT_AUTH/controllers"
	"github.com/Ritik100-AIT/Go_JWT_AUTH/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	route.Use(middleware.Authenticate())
	route.GET("users", controllers.GetUsers())
	route.GET("users/:id", controllers.GetUser())
}
