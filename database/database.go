package database

import (
	"context"
	"errors"
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
	db.AutoMigrate(&models.Cart{})
	db.AutoMigrate(&models.CartItem{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.OrderItem{})
	db.AutoMigrate(&models.CouponObject{})
	fmt.Println("Connected to sqlite...")

	// MigrateDB(db)
}

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Cart{}, &models.CartItem{}, &models.Order{}, &models.OrderItem{})
	if err != nil {
		panic("failed to migrate database")
	}
	return err
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

func AddUser(user models.User) error {
	if err := db.Create(&user).Error; err != nil {
		return fmt.Errorf("AddUser: %v", err)
	}
	return nil
}

func SaveUser(user models.User) error {
	if err := db.Save(&user).Error; err != nil {
		return fmt.Errorf("SaveUser: %v", err)
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

func GetUserCart(id string) (models.Cart, error) {
	var cart models.Cart
	err := db.Preload("Items.Product").Where("user_id = ?", id).First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return cart, fmt.Errorf("GetUserCart: no cart found for user ID %s", id)
		}
		return cart, fmt.Errorf("GetUserCart: %v", err)
	}
	return cart, nil
}

func GetUserOrders(id string) ([]models.Order, error) {
	var orders []models.Order
	err := db.Preload("OrderItems.Product").Where("user_id = ?", id).Find(&orders).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return orders, fmt.Errorf("GetUserOrders: no orders found for user ID %s", id)
		}
		return orders, fmt.Errorf("GetUserOrders: %v", err)
	}
	return orders, nil
}

func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := db.Preload("OrderItems.Product").Find(&orders).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return orders, fmt.Errorf("GetAllOrders: no orders found")
		}
		return orders, fmt.Errorf("GetAllOrders: %v", err)
	}
	return orders, nil
}

func AddItemToCart(userID string, productID uint, quantity int) error {
	var cart models.Cart
	err := db.Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create a new cart if none exists
			cart = models.Cart{UserID: userID}
			if err := db.Create(&cart).Error; err != nil {
				return fmt.Errorf("AddItemToCart: failed to create new cart: %v", err)
			}
		} else {
			return fmt.Errorf("AddItemToCart: %v", err)
		}
	}

	// Check if the product is already in the cart
	var item models.CartItem
	err = db.Where("cart_id = ? AND product_id = ?", cart.ID, productID).First(&item).Error
	if err == nil {
		// Update quantity if the item exists
		item.Quantity += quantity
		return db.Save(&item).Error
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("AddItemToCart: %v", err)
	}

	// Add new item to the cart
	newItem := models.CartItem{CartID: cart.ID, ProductID: productID, Quantity: quantity}
	return db.Create(&newItem).Error
}

func RemoveItemFromCart(userID string, productID uint, quantity int) error {
	var cart models.Cart

	// Find the cart for the user
	err := db.Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("RemoveItemFromCart: no cart found for user ID %s", userID)
		}
		return fmt.Errorf("RemoveItemFromCart: %v", err)
	}

	// Find the item in the cart
	var item models.CartItem
	err = db.Where("cart_id = ? AND product_id = ?", cart.ID, productID).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("RemoveItemFromCart: product ID %d not found in cart", productID)
		}
		return fmt.Errorf("RemoveItemFromCart: %v", err)
	}

	// Adjust quantity or remove item
	if quantity >= item.Quantity {
		// Remove the item entirely if quantity to remove is greater than or equal to current quantity
		if err := db.Delete(&item).Error; err != nil {
			return fmt.Errorf("RemoveItemFromCart: failed to remove item: %v", err)
		}
	} else {
		// Decrease the quantity
		item.Quantity -= quantity
		if err := db.Save(&item).Error; err != nil {
			return fmt.Errorf("RemoveItemFromCart: failed to update item quantity: %v", err)
		}
	}

	return nil
}

func PlaceOrder(userID string, paymentMethod string, shippingDetails string, couponCode string) error {

	// Step 1: Retrieve the user's cart
	var cart models.Cart
	err := db.Preload("Items.Product").Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("PlaceOrder: no cart found for user ID %s", userID)
		}
		return fmt.Errorf("PlaceOrder: error fetching cart: %v", err)
	}

	if len(cart.Items) == 0 {
		return fmt.Errorf("PlaceOrder: cart is empty")
	}

	// Step 2: Calculate total price and apply coupon discount
	var totalPrice float64 = 0
	for _, cartItem := range cart.Items {
		totalPrice += float64(cartItem.Quantity) * cartItem.Product.Price
	}

	var discount float64 = 0
	if couponCode != "" {
		coupons := map[string]float64{
			"SAVE10": 0.10,
			"SAVE20": 0.20,
		}
		if discountRate, exists := coupons[couponCode]; exists {
			discount = totalPrice * discountRate
			totalPrice -= discount
		}
	}

	// Step 3: Create a new order

	order := models.Order{
		UserID:          userID,
		PaymentMethod:   paymentMethod,
		Status:          "Pending",
		TotalPrice:      totalPrice,
		Discount:        discount,
		CouponCode:      couponCode,
		ShippingDetails: shippingDetails,
	}

	fmt.Println(order)

	err = db.Create(&order).Error
	if err != nil {
		return fmt.Errorf("PlaceOrder: error creating order: %v", err)
	}

	// Step 4: Add items from the cart to the order

	var orderItems []models.OrderItem
	for _, cartItem := range cart.Items {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: cartItem.ProductID,
			Product:   cartItem.Product,
			Quantity:  cartItem.Quantity,
			Price:     cartItem.Product.Price,
		}
		orderItems = append(orderItems, orderItem)
	}

	err = db.Create(&orderItems).Error
	if err != nil {
		return fmt.Errorf("PlaceOrder: error adding items to order: %v", err)
	}

	var user models.User
	err = db.Where("ID = ?", userID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("PlaceOrder: no record found for user ID %s", userID)
		}
	}
	user.OrdersCount++
	SaveUser(user)

	// Step 5: Clear the user's cart
	err = db.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error
	if err != nil {
		return fmt.Errorf("PlaceOrder: error clearing cart items: %v", err)
	}

	return nil
}

func AddCoupon(coupon models.CouponObject) error {
	if err := db.Create(&coupon).Error; err != nil {
		return fmt.Errorf("AddCoupon: %v", err)
	}
	return nil
}

func SaveCoupon(coupon models.CouponObject) error {
	if err := db.Save(&coupon).Error; err != nil {
		return fmt.Errorf("SaveCoupon: %v", err)
	}
	return nil
}

func GetAllCoupons() ([]models.CouponObject, error) {
	var coupon []models.CouponObject
	if err := db.Find(&coupon).Error; err != nil {
		return nil, fmt.Errorf("get all coupon: %v", err)
	}
	return coupon, nil
}

func GetCoupon(code string) (models.CouponObject, error) {
	var coupon models.CouponObject
	if err := db.Where("code = ?", code).First(&coupon).Error; err != nil {
		return coupon, fmt.Errorf("GetCoupon: %v", err)
	}
	return coupon, nil
}

func DeleteCoupon(coupon models.CouponObject) error {
	if err := db.Where("code = ?", coupon.Code).Delete(&coupon).Error; err != nil {
		return fmt.Errorf("DeleteCoupon: %v", err)
	}
	return nil
}
