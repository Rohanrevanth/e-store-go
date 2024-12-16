# E-Commerce Website Project

## Overview
This project is a comprehensive e-commerce platform built using a Go backend and an Angular frontend. It supports essential e-commerce functionalities such as user management, product catalog, orders, discounts, and analytics for an admin panel.

### Key Features:
- User Authentication and Authorization
- Product Listing and Categorization
- Order Management
- Coupon and Discount Handling
- Admin Dashboard with Analytics
- RESTful APIs for seamless frontend-backend communication

## Technologies Used
- **Backend**: Go with Gin framework
- **Frontend**: Angular (hosted locally at `http://localhost:4200` during development)
- **Database**: SQLite, populated with categories and products.(Added a test_backup.db with same default data if chooses to start afresh) 
- **ORM**: GORM (for database interactions)
- **API Testing**: Postman

---

## API Endpoints

### Authentication APIs
1. **Register User**
   - `POST /api/auth/register`
   - **Body**: `{ "username": "string", "email": "string", "password": "string" }`
   - **Response**: User registration status.

2. **Register Admin**
   - `POST /api/auth/login`
   - **Body**: `{ "email": "string", "password": "string", "type": "ADMIN" }`
   - **Response**: JWT Token.

3. **Login**
   - `POST /api/auth/login`
   - **Body**: `{ "email": "string", "password": "string" }`
   - **Response**: JWT Token.

---

### Product APIs
1. **Get All Products**
   - `GET /api/products`
   - **Response**: List of products.

2. **Get Product by ID**
   - `GET /api/products/:id`
   - **Response**: Product details.

3. **Add Product** (Admin only)
   - `POST /api/products`
   - **Body**: `{ "name": "string", "description": "string", "price": float, "category": "string", "image": "string" }`
   - **Response**: Status of product addition.

4. **Update Product** (Admin only)
   - `PUT /api/products/:id`
   - **Body**: `{ "name": "string", "description": "string", "price": float, "category": "string", "image": "string" }`
   - **Response**: Status of product update.

5. **Delete Product** (Admin only)
   - `DELETE /api/products/:id`
   - **Response**: Status of product deletion.

---

### Order APIs
1. **Place an Order**
   - `POST /api/orders`
   - **Response**: Order creation status.

2. **Get User Orders**
   - `GET /api/orders/user/:user_id`
   - **Response**: List of user orders.

3. **Get All Orders** (Admin only)
   - `GET /api/orders`
   - **Response**: List of all orders.

---

### Coupon APIs
1. **Get All Coupons**
   - `GET /api/coupons`
   - **Response**: List of all available coupons.

2. **Add Coupon** (Admin only)
   - `POST /api/coupons`
   - **Body**: `{ "code": "string", "discount": float, "order_frequency": int64 }`
   - **Response**: Status of coupon addition.

3. **Delete Coupon** (Admin only)
   - `DELETE /api/coupons/:id`
   - **Response**: Status of coupon deletion.

---

## Local Development Setup

### Prerequisites
- Go installed (version 1.18+ recommended)
- Node.js and npm installed (Angular CLI required)
- SQLite or any database supported by GORM

### Backend Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/Rohanrevanth/e-store-go.git
   cd e-store
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the server:
   ```bash
   go run .
   ```

   The backend will be available at `http://localhost:8080`.

### Frontend Setup
1. Navigate to the frontend directory:
   ```bash
   cd e-store-ng
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the Angular app:
   ```bash
   ng serve
   ```

   The frontend will be available at `http://localhost:4200`.

---

## Deployment

### Backend Deployment
- Use Docker for containerization.
- Deploy to any cloud provider (e.g., AWS, GCP, Azure) or on-premises server.

### Frontend Deployment
- Build the Angular project:
  ```bash
  ng build --prod
  ```
- Deploy the `dist/` folder to a static hosting service (e.g., Netlify, Vercel, or AWS S3).

---

## Future Enhancements
- Payment Gateway Integration (e.g., Stripe, PayPal)
- Real-Time Notifications
- Wishlist and Cart Functionality
- Enhanced Product Recommendations
- Multi-language Support

---

## Contributors
- [Your Name](https://github.com/your-profile)

---

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

