package routes

import (
	"github.com/Rohanrevanth/e-store-go/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	//User routes
	router.POST("/login", controllers.Login)
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/user/:id", controllers.GetUserByID)
	router.POST("/register", controllers.RegisterUsers)
	router.POST("/delete", controllers.DeleteUser)

	// protected := router.Group("/").Use(auth.JWTAuthMiddleware())
	// {
	// 	protected.GET("/users", controllers.GetAllUsers)
	// }
}
