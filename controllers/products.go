package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Rohanrevanth/e-store-go/database"
	"github.com/Rohanrevanth/e-store-go/models"
	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	categories, err := database.GetAllCategories()
	if err != nil {
		log.Println("Error fetching categories:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": categories})
}

func GetBestSellers(c *gin.Context) {
	bestSellers, err := database.GetBestSellers()
	if err != nil {
		log.Println("Error fetching best-sellers:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to fetch best-sellers"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": bestSellers})
}

func GetProducts(c *gin.Context) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		log.Println("Error binding body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request payload"})
		return
	}
	products, err := database.GetProducts(product.Category)
	if err != nil {
		log.Println("Error fetching products:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": products})
}

func GetAllProducts(c *gin.Context) {
	products, err := database.GetAllProducts()
	if err != nil {
		log.Println("Error fetching products:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": products})
}

func AddCategories(c *gin.Context) {
	var newCategories []models.Category
	if err := c.BindJSON(&newCategories); err != nil {
		log.Println("Error binding categories:", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request payload"})
		return
	}

	// var addedCategories []models.Category
	for _, category := range newCategories {
		err := database.AddCategory(category)
		if err != nil {
			log.Println("Error adding category:", err)
			continue // Skip this category and proceed with the others
		}

		// savedUser, err := database.GetUserByEmail(user.Email)
		// if err != nil {
		// 	log.Println("Error fetching saved user:", err)
		// 	continue
		// }
		// registeredUsers = append(registeredUsers, savedUser)
	}

	// c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Categories added", "data": registeredUsers})
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Categories added"})
}

func AddProducts(c *gin.Context) {
	var newProducts []models.Product
	if err := c.BindJSON(&newProducts); err != nil {
		log.Println("Error binding product:", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request payload"})
		return
	}

	for _, product := range newProducts {
		detailsJSON, err := json.Marshal(product.Details)
		if err != nil {
			log.Fatalf("Error marshaling details: %v", err)
		}
		product.Details = string(detailsJSON)

		err = database.AddProduct(product)
		if err != nil {
			log.Println("Error adding product:", err)
			continue // Skip this product and proceed with the others
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Products added"})
}

func GetUserByID_(c *gin.Context) {
	id := c.Param("id")
	// ctx := context.Background()

	// Attempt to retrieve the user from Redis
	// cachedUser, err := database.RedisClient.Get(ctx, id).Result()
	// if err == nil {
	// 	// Cache hit
	// 	var user models.User
	// 	if jsonErr := json.Unmarshal([]byte(cachedUser), &user); jsonErr == nil {
	// 		c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
	// 		return
	// 	}
	// }

	// Cache miss: retrieve from database
	user, err := database.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
		return
	}

	// Cache the user in Redis
	// userJSON, _ := json.Marshal(user)
	// if err := database.RedisClient.Set(ctx, id, userJSON, 10*time.Minute).Err(); err != nil {
	// 	log.Println("Failed to cache user:", err)
	// }

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

func DeleteUser_(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind user"})
		return
	}
	err := database.DeleteUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func RegisterUsers_(c *gin.Context) {
	var newUsers []models.User
	if err := c.BindJSON(&newUsers); err != nil {
		log.Println("Error binding users:", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request payload"})
		return
	}

	var registeredUsers []models.User
	for _, user := range newUsers {
		// Validate user fields here (e.g., Email and Password)

		if err := user.HashPassword(user.Password); err != nil {
			log.Println("Error hashing password:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to hash password"})
			return
		}

		err := database.SignupUser(user)
		if err != nil {
			log.Println("Error registering user:", err)
			continue // Skip this user and proceed with the others
		}

		savedUser, err := database.GetUserByEmail(user.Email)
		if err != nil {
			log.Println("Error fetching saved user:", err)
			continue
		}
		registeredUsers = append(registeredUsers, savedUser)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Users processed", "data": registeredUsers})
}
