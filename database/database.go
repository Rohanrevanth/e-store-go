package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Rohanrevanth/e-store-go/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
)

var db *gorm.DB
var RedisClient *redis.Client

func ConnectDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database!", err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Product{})
	fmt.Println("Connected to sqlite...")
}

func InitializeRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password by default
		DB:       0,                // Default DB
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully")
}

func GetUserByEmail(email string) (models.User, error) {
	var usr models.User
	if err := db.Where("email = ?", email).First(&usr).Error; err != nil {
		return usr, fmt.Errorf("GetUserByEmail: %v", err)
	}
	return usr, nil
}

func GetUserByID(id string) (models.User, error) {
	var usr models.User
	if err := db.Where("ID = ?", id).First(&usr).Error; err != nil {
		return usr, fmt.Errorf("GetUserByID: %v", err)
	}
	return usr, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("get all users: %v", err)
	}
	return users, nil
}

func SignupUser(user models.User) error {
	if err := db.Create(&user).Error; err != nil {
		return fmt.Errorf("SignupUser: %v", err)
	}
	return nil
}

func DeleteUser(user models.User) error {
	if err := db.Delete(&user).Error; err != nil {
		return fmt.Errorf("DeleteUser: %v", err)
	}
	return nil
}

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		return nil, fmt.Errorf("get all categories: %v", err)
	}
	return categories, nil
}

func GetBestSellers() ([]models.Product, error) {
	var products []models.Product
	if err := db.Where("isbestseller = ?", true).Find(&products).Error; err != nil {
		return nil, fmt.Errorf("get all products: %v", err)
	}
	return products, nil
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := db.Find(&products).Error; err != nil {
		return nil, fmt.Errorf("get all products: %v", err)
	}
	return products, nil
}

func GetProducts(category string) ([]models.Product, error) {
	var products []models.Product
	if err := db.Where("category = ?", category).Find(&products).Error; err != nil {
		return products, fmt.Errorf("GetUserByEmail: %v", err)
	}
	return products, nil
}

func AddCategory(category models.Category) error {
	if err := db.Create(&category).Error; err != nil {
		return fmt.Errorf("AddCategory: %v", err)
	}
	return nil
}

func AddProduct(product models.Product) error {
	if err := db.Create(&product).Error; err != nil {
		return fmt.Errorf("AddProduct: %v", err)
	}
	return nil
}
