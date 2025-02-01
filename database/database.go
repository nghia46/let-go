package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	// Tải biến môi trường từ file .env
	if err := godotenv.Load(); err != nil {
		log.Printf("Không thể tải file .env: %v", err)
	} else {
		log.Println("File .env đã được tải thành công")
	}

	// Kiểm tra giá trị của các biến môi trường
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))

	// Tạo chuỗi kết nối từ biến môi trường
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require pool_mode= session",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "Abcd1234@"),
		getEnv("DB_NAME", "postgres"),
		getEnv("DB_PORT", "5432"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Không thể kết nối database:", err)
	} else {
		log.Println("Kết nối đến database thành công")
	}

}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
