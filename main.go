package main

import (
	"let-go/database"
	"let-go/handlers"
	"let-go/models"

	_ "let-go/docs" // Thêm dòng này để load Swagger docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {
	database.ConnectDatabase()

	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Product{})

	r := gin.Default()

	// Cung cấp file tĩnh cho CSS
	r.Static("/assets", "./assets")

	// Đường dẫn API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Group routes cho Users
	users := r.Group("/users")
	{
		users.GET("/", handlers.GetUsers)    // GET /users
		users.POST("/", handlers.CreateUser) // POST /users
	}

	// Group routes cho Products
	products := r.Group("/products")
	{
		products.GET("/", handlers.GetProducts)    // GET /products
		products.POST("/", handlers.CreateProduct) // POST /products
	}

	r.Run(":8080")
}
