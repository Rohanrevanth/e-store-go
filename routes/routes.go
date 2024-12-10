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

	//Product routes
	router.GET("/categories", controllers.GetAllCategories)
	router.POST("/categories", controllers.AddCategories)
	router.GET("/best-sellers", controllers.GetBestSellers)
	router.GET("/all-products", controllers.GetAllProducts)
	router.POST("/add-products", controllers.AddProducts)
	router.POST("/get-products", controllers.GetProducts)

	// protected := router.Group("/").Use(auth.JWTAuthMiddleware())
	// {
	// 	protected.GET("/users", controllers.GetAllUsers)
	// }
}
