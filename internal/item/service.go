package item

import (
	"backend/internal/constant"
	"backend/internal/model"

	"gorm.io/gorm"
)

// Service struct ใช้สำหรับติดต่อกับ Repository เพื่อทำงานกับฐานข้อมูลผ่าน Repository
type Service struct {
	Repository *Repository // เชื่อมกับ Repository เพื่อใช้เรียกฟังก์ชันที่เกี่ยวข้องกับฐานข้อมูล
}

// NewService สร้าง Service ใหม่ โดยรับตัวแปร db ของ GORM เพื่อสร้าง Repository
func NewService(db *gorm.DB) Service {
	return Service{
		Repository: NewRepository(db), // คืนค่าเป็น Service ใหม่พร้อม Repository ที่เชื่อมต่อกับ db
	}
}

// CreateItem สร้างรายการสินค้าใหม่ โดยรับ request และ userId จากผู้ใช้งาน
func (s Service) CreateItem(req model.RequestItem, userId uint) (model.Item, error) {
	// กำหนดค่าของสินค้าจาก request และ userId
	item := model.Item{
		Title:    req.Title,
		Price:    req.Price,
		Quantity: req.Quantity,
		Status:   constant.ItemPendingStatus, // ตั้งค่า Status เป็น Pending โดยค่าเริ่มต้น
		OwnerID:  userId,                     // กำหนดเจ้าของสินค้า
	}

	// เรียกใช้ Repository เพื่อบันทึกสินค้าใหม่ลงในฐานข้อมูล
	id, err := s.Repository.CreateItem(&item)
	if err != nil {
		return model.Item{}, err // หากเกิดข้อผิดพลาด ส่งกลับ error
	}

	item.ID = id // กำหนด ID ที่ได้จากฐานข้อมูลให้กับ item
	return item, nil // ส่งกลับสินค้าใหม่ที่สร้าง
}

// GetItems ดึงข้อมูลรายการสินค้าทั้งหมด โดยเรียกใช้ Repository
func (s Service) GetItems() ([]model.Item, error) {
	// เรียกใช้ Repository เพื่อดึงสินค้าทั้งหมด
	items, err := s.Repository.GetItems()
	if err != nil {
		return nil, err // หากเกิดข้อผิดพลาด ส่งกลับ error
	}
	return items, nil // ส่งกลับรายการสินค้าที่ดึงมาได้
}

// DeleteItem ลบสินค้าตาม itemId ที่ระบุ
func (s Service) DeleteItem(itemId uint) error {
	// เรียกใช้ Repository เพื่อทำการลบสินค้า
	err := s.Repository.DeleteItem(itemId)
	if err != nil {
		return err // หากเกิดข้อผิดพลาด ส่งกลับ error
	}
	return nil // ส่งกลับ nil เมื่อการลบสำเร็จ
}

// UpdateItem อัปเดตรายการสินค้าตาม itemId และข้อมูลใหม่ที่ระบุใน req
func (s Service) UpdateItem(itemId uint, req model.RequestItem) (model.Item, error) {
	// ดึงสินค้าปัจจุบันจากฐานข้อมูล
	item, err := s.Repository.GetItemByID(itemId)
	if err != nil {
		return model.Item{}, err // หากเกิดข้อผิดพลาด ส่งกลับ error
	}

	// อัปเดตข้อมูลของสินค้าตามค่าที่ได้รับจาก req
	item.Title = req.Title
	item.Price = req.Price
	item.Quantity = req.Quantity

	// เรียกใช้ Repository เพื่อบันทึกการอัปเดตสินค้า
	updatedItem, err := s.Repository.UpdateItem(item)
	if err != nil {
		return model.Item{}, err // หากเกิดข้อผิดพลาด ส่งกลับ error
	}

	return updatedItem, nil // ส่งกลับข้อมูลสินค้าที่ถูกอัปเดต
}

// GetItem ดึงข้อมูลรายการสินค้าตาม itemId โดยเรียกใช้ Repository
func (s Service) GetItem(itemId uint) (model.Item, error) {
	// เรียกใช้ Repository เพื่อดึงข้อมูลสินค้าตาม ID ที่ระบุ
	item, err := s.Repository.GetItemByID(itemId)
	if err != nil {
		return model.Item{}, err // หากเกิดข้อผิดพลาด ส่งกลับ error
	}
	return item, nil // ส่งกลับข้อมูลสินค้าที่ดึงได้
}
