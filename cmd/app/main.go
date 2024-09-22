package main

import (
	"backend/internal/item"
	"backend/internal/model"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("POSTGRES_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if err := db.AutoMigrate(&model.Item{}); err != nil {
		log.Fatal("Failed to auto-migrate models:", err)
	}

	r := gin.Default()

	FRONTEND_URL := os.Getenv("FRONTEND_URL")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{FRONTEND_URL}, // ใช้ URL ที่ได้จาก .env
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	itemController := item.NewController(db)

	
	r.POST("/items", itemController.CreateItem)
	r.GET("/items", itemController.GetItems)
	r.GET("/items/:id", itemController.GetItem)
	r.DELETE("/items/:id", itemController.DeleteItem)
	r.PUT("/items/:id", itemController.PutItem)

	// ตั้งค่า PORT ที่ใช้ในการรันเซิร์ฟเวอร์
	port := os.Getenv("PORT_API")


	if err := r.Run(":" + port); err != nil {
		log.Panic("Failed to start server:", err)
	}
}
