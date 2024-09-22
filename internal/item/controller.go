package item

import (
	"backend/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Controller struct ประกอบด้วย Service สำหรับการจัดการข้อมูลสินค้า
type Controller struct {
	service Service
}

// NewController ฟังก์ชันสำหรับสร้าง Controller ใหม่ โดยรับ db จาก GORM มาเชื่อมต่อกับ Service
func NewController(db *gorm.DB) *Controller {
	return &Controller{service: NewService(db)} // สร้างและเชื่อมต่อ Service ใหม่
}

// CreateItem ฟังก์ชันสำหรับเพิ่มสินค้าลงในฐานข้อมูล
func (c *Controller) CreateItem(ctx *gin.Context) {
	var request model.RequestItem // ประกาศตัวแปรเพื่อรับข้อมูลจากผู้ใช้
	if err := ctx.ShouldBindJSON(&request); err != nil { // ตรวจสอบและบันทึกข้อมูลที่ส่งมาในรูปแบบ JSON
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // ส่งข้อความแสดงข้อผิดพลาดหากข้อมูลไม่ถูกต้อง
		return
	}

	// สมมติให้ userId มาจากการตรวจสอบผู้ใช้ (เช่นจากระบบ authentication)
	userId := 1 // เปลี่ยนให้เหมาะสมตามโปรเจคจริง

	item, err := c.service.CreateItem(request, uint(userId)) // เรียกใช้ service เพื่อสร้างสินค้าจากข้อมูลที่ได้รับ
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // แสดงข้อผิดพลาดหากไม่สามารถสร้างสินค้าได้
		return
	}

	ctx.JSON(http.StatusCreated, item) // ส่งข้อมูลสินค้าที่ถูกสร้างกลับไปยังผู้ใช้
}

// GetItems ฟังก์ชันสำหรับดึงข้อมูลสินค้าทั้งหมดจากฐานข้อมูล
func (c *Controller) GetItems(ctx *gin.Context) {
	items, err := c.service.GetItems() // เรียกใช้ service เพื่อดึงข้อมูลสินค้าทั้งหมด
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // แสดงข้อผิดพลาดหากไม่สามารถดึงข้อมูลได้
		return
	}
	ctx.JSON(http.StatusOK, items) // ส่งข้อมูลสินค้ากลับไปยังผู้ใช้ในรูปแบบ JSON
}

// DeleteItem ฟังก์ชันสำหรับลบสินค้าตาม ID
func (c *Controller) DeleteItem(ctx *gin.Context) {
	// ดึง item ID จากพารามิเตอร์ของ URL
	itemIdParam := ctx.Param("id")

	// แปลง itemIdParam จาก string เป็น uint
	idUint, err := strconv.ParseUint(itemIdParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID format"}) // ส่งข้อความแสดงข้อผิดพลาดหาก ID ไม่ถูกต้อง
		return
	}
	itemId := uint(idUint)

	// เรียกใช้ service เพื่อลบสินค้า
	err = c.service.DeleteItem(itemId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // แสดงข้อผิดพลาดหากไม่สามารถลบได้
		return
	}

	// ส่งข้อความยืนยันว่าลบสินค้าเสร็จสิ้น
	ctx.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}

// PutItem ฟังก์ชันสำหรับอัปเดตข้อมูลสินค้าตาม ID
func (c *Controller) PutItem(ctx *gin.Context) {
	// ดึง item ID จากพารามิเตอร์ของ URL
	itemIdParam := ctx.Param("id")

	// แปลง itemIdParam เป็น uint
	idUint, err := strconv.ParseUint(itemIdParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID format"}) // ส่งข้อความแสดงข้อผิดพลาดหาก ID ไม่ถูกต้อง
		return
	}
	itemId := uint(idUint)

	// ผูกข้อมูล JSON ที่ส่งมาจากผู้ใช้กับ struct ของ model.RequestItem
	var request model.RequestItem
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // แสดงข้อผิดพลาดหาก JSON ไม่ถูกต้อง
		return
	}

	// เรียกใช้ service เพื่ออัปเดตสินค้า
	updatedItem, err := c.service.UpdateItem(itemId, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // แสดงข้อผิดพลาดหากไม่สามารถอัปเดตได้
		return
	}

	// ส่งข้อมูลสินค้าที่ถูกอัปเดตกลับไปยังผู้ใช้
	ctx.JSON(http.StatusOK, updatedItem)
}

// GetItem ฟังก์ชันสำหรับดึงข้อมูลสินค้าตาม ID
func (c *Controller) GetItem(ctx *gin.Context) {
	// ดึง item ID จากพารามิเตอร์ของ URL
	itemIdParam := ctx.Param("id")

	// แปลง itemIdParam เป็น uint
	idUint, err := strconv.ParseUint(itemIdParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID format"}) // แสดงข้อผิดพลาดหาก ID ไม่ถูกต้อง
		return
	}
	itemId := uint(idUint)

	// เรียกใช้ service เพื่อนำข้อมูลสินค้ามาแสดง
	item, err := c.service.GetItem(itemId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // แสดงข้อผิดพลาดหากไม่สามารถดึงข้อมูลได้
		return
	}

	// ส่งข้อมูลสินค้าที่ดึงมาให้ผู้ใช้
	ctx.JSON(http.StatusOK, item)
}
