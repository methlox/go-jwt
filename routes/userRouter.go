package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/methlox/go-jwt/controllers"
	"github.com/methlox/go-jwt/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
