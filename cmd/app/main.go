package main

import (
	"backend/internal/item"     // Import package item ที่เป็นส่วนจัดการสินค้าภายในระบบ
	"backend/internal/model"    // Import package model ที่มีการกำหนดโมเดลของฐานข้อมูล
	"log"
	"os"

	"github.com/gin-contrib/cors"  // Import middleware สำหรับการจัดการ Cross-Origin Resource Sharing (CORS)
	"github.com/gin-gonic/gin"     // Import Gin web framework สำหรับสร้าง API
	"github.com/joho/godotenv"     // Import godotenv สำหรับโหลดไฟล์ .env ที่มีค่าคอนฟิกต่างๆ
	"gorm.io/driver/postgres"      // Import driver ของ GORM สำหรับการเชื่อมต่อกับฐานข้อมูล PostgreSQL
	"gorm.io/gorm"                 // Import GORM ORM สำหรับการจัดการฐานข้อมูล
)

func main() {

	// โหลดไฟล์ .env ที่เก็บค่าคอนฟิกต่างๆ เช่น POSTGRES_DSN และ FRONTEND_URL
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")  // ถ้าโหลดไฟล์ .env ไม่ได้ ให้หยุดการทำงานและแจ้ง error
	}

	// ดึงค่าตัวแปร POSTGRES_DSN จากไฟล์ .env เพื่อนำไปใช้เชื่อมต่อกับฐานข้อมูล
	dsn := os.Getenv("POSTGRES_DSN")
	
	// เปิดการเชื่อมต่อกับฐานข้อมูล PostgreSQL ผ่าน GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)  // ถ้าเชื่อมต่อฐานข้อมูลไม่สำเร็จ ให้หยุดการทำงานและแจ้ง error
	}

	// Auto-migrate ทำการสร้างหรืออัปเดตโครงสร้างของตารางในฐานข้อมูลตามโมเดลที่กำหนด
	if err := db.AutoMigrate(&model.Item{}); err != nil {
		log.Fatal("Failed to auto-migrate models:", err)  // ถ้าไม่สามารถ migrate โมเดลได้ ให้หยุดการทำงานและแจ้ง error
	}

	// เริ่มต้นการใช้งาน Gin framework โดยใช้การตั้งค่าเริ่มต้น
	r := gin.Default()

	// ดึงค่าตัวแปร FRONTEND_URL จากไฟล์ .env สำหรับการตั้งค่า CORS
	FRONTEND_URL := os.Getenv("FRONTEND_URL")

	// ตั้งค่า CORS เพื่ออนุญาตให้มีการเชื่อมต่อมาจาก URL ที่ระบุใน .env และกำหนดวิธีการเชื่อมต่อ (GET, POST, PUT, DELETE)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{FRONTEND_URL}, // ใช้ URL ที่ได้จาก .env สำหรับอนุญาตการเชื่อมต่อ
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"}, // อนุญาตเฉพาะวิธีการเชื่อมต่อเหล่านี้
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // อนุญาตเฉพาะ headers เหล่านี้
		AllowCredentials: true, // อนุญาตให้ส่ง cookies หรือข้อมูล credential ผ่าน CORS ได้
	}))

	// สร้าง controller สำหรับจัดการสินค้า โดยส่ง database ที่เชื่อมต่อเข้าไปใน controller
	itemController := item.NewController(db)

	// กำหนดเส้นทาง API สำหรับจัดการรายการสินค้า
	r.POST("/items", itemController.CreateItem)   // สร้างสินค้าใหม่
	r.GET("/items", itemController.GetItems)      // ดึงข้อมูลสินค้าทั้งหมด
	r.GET("/items/:id", itemController.GetItem)   // ดึงข้อมูลสินค้าตาม ID
	r.DELETE("/items/:id", itemController.DeleteItem) // ลบสินค้าตาม ID
	r.PUT("/items/:id", itemController.PutItem)   // อัปเดตข้อมูลสินค้าตาม ID

	// ตั้งค่า PORT ที่ใช้ในการรันเซิร์ฟเวอร์ โดยดึงค่าจากไฟล์ .env
	port := os.Getenv("PORT_API")

	// รันเซิร์ฟเวอร์บนพอร์ตที่กำหนด ถ้าเกิดข้อผิดพลาดจะหยุดการทำงานและแจ้ง error
	if err := r.Run(":" + port); err != nil {
		log.Panic("Failed to start server:", err)
	}
}
