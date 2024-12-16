package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a user model in the database
type User struct {
	gorm.Model
	Username     string    `json:"username" gorm:"unique"`
	Email        string    `json:"email" gorm:"unique"`
	Password     string    `json:"password"`
	OrdersCount  int64     `json:"orders_count,omitempty"`
	SavedAddress Addresses `json:"saved_address,omitempty" gorm:"type:json"`
	Type         string    `json:"type,omitempty"`
}

type Addresses []string

type AddressStringObj struct {
	Address Addresses `json:"address"`
}

type CouponCodeObj struct {
	CouponCode string `json:"coupon_code"`
}

type CouponObject struct {
	gorm.Model
	Code           string  `json:"code"`
	Discount       float64 `json:"discount"`
	OrderFrequency int64   `json:"order_frequency"`
}

type Cart struct {
	gorm.Model
	UserID string     `json:"user_id" gorm:"unique"`          // Each cart belongs to a specific user
	User   User       `json:"user" gorm:"foreignKey:UserID"`  // Foreign key for User
	Items  []CartItem `json:"items" gorm:"foreignKey:CartID"` // Establishes a relationship with CartItem
}

type CartItem struct {
	gorm.Model
	CartID    uint    `json:"cart_id"`                             // Foreign key to associate with Cart
	ProductID uint    `json:"product_id"`                          // Foreign key to associate with Product
	Product   Product `json:"product" gorm:"foreignKey:ProductID"` // Reference to the Product
	Quantity  int     `json:"quantity"`                            // Quantity of the product in the cart
}

type Order struct {
	gorm.Model
	UserID          string      `json:"user_id" gorm:"not null"`
	PaymentMethod   string      `json:"payment_method,omitempty"`
	Status          string      `json:"status,omitempty" gorm:"default:Pending"`
	OrderItems      []OrderItem `json:"order_items,omitempty" gorm:"foreignKey:OrderID"`
	TotalPrice      float64     `json:"total_price,omitempty"`
	Discount        float64     `json:"discount,omitempty"`
	CouponCode      string      `json:"coupon_code,omitempty"`
	ShippingDetails string      `json:"shipping_details,omitempty"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id" gorm:"not null"`            // ForeignKey to Order
	ProductID uint    `json:"product_id" gorm:"not null"`          // ForeignKey to Product
	Product   Product `json:"product" gorm:"foreignKey:ProductID"` // Product reference
	Quantity  int     `json:"quantity" gorm:"not null"`
	Price     float64 `json:"price" gorm:"not null"`
}

// type ShippingDetails struct {
// 	Name    string `json:"name" gorm:"not null"`
// 	Address string `json:"address" gorm:"not null"`
// 	City    string `json:"city" gorm:"not null"`
// 	Zip     string `json:"zip" gorm:"not null"`
// 	Phone   string `json:"phone" gorm:"not null"` // Make sure Phone is exported
// }

func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// Implement the `Value` method for saving into the database
func (a Addresses) Value() (driver.Value, error) {
	return json.Marshal(a) // Convert to JSON string
}

// Implement the `Scan` method for reading from the database
func (a *Addresses) Scan(value interface{}) error {
	var byteValue []byte

	switch v := value.(type) {
	case string:
		byteValue = []byte(v) // Convert string to []byte
	case []byte:
		byteValue = v // Already a []byte
	default:
		return errors.New("unsupported data type for SavedAddress")
	}

	return json.Unmarshal(byteValue, a) // Parse JSON into Addresses
}
