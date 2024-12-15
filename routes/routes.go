package routes

import (
	"github.com/Rohanrevanth/e-store-go/auth"
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

	router.GET("/get-orders/:id", controllers.GetUserOders)

	router.GET("/get-coupons", controllers.GetCoupons)
	// router.GET("/get-coupon/:id", controllers.SaveCoupon)
	router.POST("/add-coupon", controllers.AddCoupon)
	router.POST("/update-coupon", controllers.SaveCoupon)
	router.POST("/delete-coupon", controllers.DeleteCoupon)

	protected := router.Group("/").Use(auth.JWTAuthMiddleware())
	{
		protected.GET("/get-cart/:id", controllers.GetUserCart)
		protected.POST("/add-to-cart/:id", controllers.AddProductToCart)
		protected.POST("/delete-from-cart/:id", controllers.RemoveItemFromCart)
		protected.POST("/save-address/:id", controllers.SaveAddress)
		protected.POST("/apply-coupon/:id", controllers.ApplyCoupon)
		protected.POST("/place-order", controllers.PlaceOrder)
	}
}
