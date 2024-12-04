package http

import (
	"time"

	"github.com/Rohanrevanth/e-store-go/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouter initializes the Gin router and registers routes.
func InitRouter() *gin.Engine {
	router := gin.Default()

	// CORS middleware configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,           // Allow cookies or authentication headers
		MaxAge:           24 * time.Hour, // Cache preflight request for 24 hours
	}))

	// Register application routes
	routes.RegisterRoutes(router)

	return router
}

// StartServer starts the HTTP server on the specified address.
func StartServer() {
	router := InitRouter()

	// Start the server on localhost:8080
	if err := router.Run("localhost:8080"); err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}
