package routes

import (
	"github.com/Rohanrevanth/e-store-go/auth"
	"github.com/Rohanrevanth/e-store-go/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.RegisterUsers)

	protected := router.Group("/").Use(auth.JWTAuthMiddleware())
	{

		//User routes
		protected.GET("/users", controllers.GetAllUsers)
		protected.GET("/user/:id", controllers.GetUserByID)
		protected.POST("/delete", controllers.DeleteUser)

		//Product routes
		protected.GET("/categories", controllers.GetAllCategories)
		protected.POST("/categories", controllers.AddCategories)
		protected.GET("/best-sellers", controllers.GetBestSellers)
		protected.GET("/all-products", controllers.GetAllProducts)
		protected.POST("/add-products", controllers.AddProducts)
		protected.POST("/get-products", controllers.GetProducts)

		//Coupon routes
		protected.GET("/get-coupons", controllers.GetCoupons)
		protected.POST("/add-coupon", controllers.AddCoupon)
		protected.POST("/update-coupon", controllers.SaveCoupon)
		protected.POST("/delete-coupon", controllers.DeleteCoupon)
		protected.POST("/apply-coupon/:id", controllers.ApplyCoupon)

		//Order routes
		protected.GET("/get-cart/:id", controllers.GetUserCart)
		protected.POST("/add-to-cart/:id", controllers.AddProductToCart)
		protected.POST("/delete-from-cart/:id", controllers.RemoveItemFromCart)
		protected.POST("/save-address/:id", controllers.SaveAddress)
		protected.GET("/get-orders/:id", controllers.GetUserOders)
		protected.GET("/get-orders", controllers.GetAllOders)
		protected.POST("/place-order", controllers.PlaceOrder)
	}
}
