package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "hieuhoccode"
	DB_PASSWORD = "hieuhoccode"
	DB_NAME     = "erp"
)

func main() {

	//---------------------------------database-------------------
	// Tạo chuỗi kết nối đến PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	// Kết nối đến cơ sở dữ liệu PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	// Tạo struct để lưu thông tin của một bản ghi trong bảng "users"
	type User struct {
		Id       int64  `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Tạo bảng "users" trong cơ sở dữ liệu
	db.AutoMigrate(&User{})

	//------------------------------API--------------------
	// Khởi tạo router sử dụng Gin
	router := gin.Default()

	// Định nghĩa route GET "/hello" để trả về "Hello, World!"
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// viết API insert user
	router.POST("/users", func(c *gin.Context) {
		// Lấy thông tin từ request

		userRequest := struct {
			Username string `json:"user_name"`
			Password string `json:"password"`
		}{}
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid request",
			})
			return
		}

		// Tạo một bản ghi mới trong bảng "users"
		user := User{Username: userRequest.Username, Password: userRequest.Password}
		db.Create(&user)

		// Trả về thông tin của bản ghi vừa được tạo
		c.JSON(200, user)
	})

	// viết API lấy danh sách user
	router.GET("/users", func(c *gin.Context) {
		// Lấy danh sách bản ghi trong bảng "users"
		var users []User
		db.Find(&users)

		// Trả về danh sách bản ghi
		c.JSON(200, users)
	})

	// viết API đăng nhập
	router.POST("/login", func(c *gin.Context) {
		// Lấy thông tin từ request
		userRequest := struct {
			Username string `json:"user_name"`
			Password string `json:"password"`
		}{}
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid request",
			})
			return
		}

		// Kiểm tra thông tin đăng nhập có hợp lệ hay không
		if userRequest.Username == "" || userRequest.Password == "" {
			c.JSON(400, gin.H{
				"message": "Invalid username or password",
			})
			return
		}

		// Tìm kiếm bản ghi trong bảng "users"từ request
		var user User
		db.Where("username = ?", userRequest.Username).First(&user)

		// Kiểm tra mật khẩu có đúng hay không
		if user.Password != userRequest.Password {
			c.JSON(401, gin.H{
				"message": "wrong username or password",
			})
			return
		}

		// Trả về thông tin của bản ghi
		c.JSON(200, user)
	})

	// Chạy server API trên cổng 8000
	router.Run(":8000")
}
