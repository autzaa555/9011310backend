package item

import (
	"backend/internal/model"

	"gorm.io/gorm"
)

// Repository struct ใช้สำหรับติดต่อกับฐานข้อมูล โดยมีตัวแปร db สำหรับการเชื่อมต่อกับ GORM
type Repository struct {
	db *gorm.DB
}

// NewRepository สร้าง Repository ใหม่ โดยรับตัวแปร db ของ GORM เพื่อการใช้งานกับฐานข้อมูล
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db} // คืนค่าเป็น Repository ใหม่
}

// CreateItem เพิ่มข้อมูลสินค้าลงในฐานข้อมูล
func (r *Repository) CreateItem(item *model.Item) (uint, error) {
	// ใช้คำสั่ง db.Create เพื่อเพิ่มรายการสินค้าใหม่ในฐานข้อมูล
	if err := r.db.Create(item).Error; err != nil {
		return 0, err // หากเกิดข้อผิดพลาด ส่งกลับค่า error
	}
	return item.ID, nil // ส่งกลับค่า ID ของสินค้าที่สร้างสำเร็จ
}

// GetItems ดึงข้อมูลสินค้าทั้งหมดจากฐานข้อมูล
func (r *Repository) GetItems() ([]model.Item, error) {
	var items []model.Item
	// ใช้คำสั่ง db.Find เพื่อดึงรายการสินค้าทั้งหมด
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err // หากเกิดข้อผิดพลาด ส่งกลับค่า error
	}
	return items, nil // ส่งกลับรายการสินค้าทั้งหมด
}

// DeleteItem ลบข้อมูลสินค้าตาม itemId
func (r *Repository) DeleteItem(itemId uint) error {
	// ใช้คำสั่ง db.Delete เพื่อลบสินค้าจากฐานข้อมูลตาม itemId ที่ได้รับ
	if err := r.db.Delete(&model.Item{}, itemId).Error; err != nil {
		return err // หากเกิดข้อผิดพลาด ส่งกลับค่า error
	}
	return nil // ส่งกลับ nil เมื่อการลบสำเร็จ
}

// GetItemByID ค้นหาสินค้าตาม ID ที่ระบุ
func (r *Repository) GetItemByID(itemId uint) (model.Item, error) {
	var item model.Item
	// ใช้คำสั่ง db.First เพื่อค้นหาสินค้าโดยใช้ itemId เป็นตัวระบุ
	if err := r.db.First(&item, itemId).Error; err != nil {
		return model.Item{}, err // หากไม่พบสินค้าหรือเกิดข้อผิดพลาด ส่งกลับค่า error
	}
	return item, nil // ส่งกลับข้อมูลสินค้าที่พบ
}

// UpdateItem อัปเดตข้อมูลสินค้าลงในฐานข้อมูล
func (r *Repository) UpdateItem(item model.Item) (model.Item, error) {
	// ใช้คำสั่ง db.Save เพื่อบันทึกการอัปเดตสินค้าลงในฐานข้อมูล
	if err := r.db.Save(&item).Error; err != nil {
		return model.Item{}, err // หากเกิดข้อผิดพลาด ส่งกลับค่า error
	}
	return item, nil // ส่งกลับข้อมูลสินค้าที่ถูกอัปเดต
}


