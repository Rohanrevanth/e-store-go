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
