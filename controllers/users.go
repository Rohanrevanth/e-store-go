package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Rohanrevanth/e-store-go/auth"
	"github.com/Rohanrevanth/e-store-go/database"
	"github.com/Rohanrevanth/e-store-go/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, err := database.GetAllUsers()
	if err != nil {
		log.Println("Error fetching users:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": users})
}

func GetUserByID(c *gin.Context) {
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

func DeleteUser(c *gin.Context) {
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

func RegisterUsers(c *gin.Context) {
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

// Login authenticates a user and returns a JWT token
func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	user, err := database.GetUserByEmail(input.Email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	input_, _ := json.Marshal(input)
	fmt.Println(string(input_))
	user_, _ := json.Marshal(user)
	fmt.Println(string(user_))
	// Check if the password is correct
	if err := user.CheckPassword(input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}
