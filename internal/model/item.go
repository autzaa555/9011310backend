package model

import "backend/internal/constant"

// โครงสร้างข้อมูล (Model) ที่แทนสินค้าภายในระบบ
type Item struct {
	ID       uint                `json:"id" gorm:"primaryKey"` // ฟิลด์ ID เป็น primary key ในฐานข้อมูลและจะถูกส่งกลับในรูปแบบ JSON
	Title    string              `json:"title"`                // ฟิลด์ Title สำหรับชื่อของสินค้า
	Price    float64             `json:"price"`                // ฟิลด์ Price สำหรับราคาของสินค้า เป็น float
	Quantity int                 `json:"quantity"`             // ฟิลด์ Quantity สำหรับจำนวนสินค้าที่มีอยู่
	Status   constant.ItemStatus `json:"status" gorm:"default:'PENDING'"` // ฟิลด์ Status สำหรับสถานะของสินค้า ใช้ enum จาก constant และค่าเริ่มต้นคือ 'PENDING'
	OwnerID  uint                `json:"owner_id"`             // ฟิลด์ OwnerID สำหรับบันทึกว่าใครเป็นเจ้าของสินค้า โดยเป็น foreign key
}

// โครงสร้างข้อมูลที่ใช้เมื่อรับคำขอสำหรับการสร้างสินค้าใหม่
type RequestItem struct {
	Title    string  `json:"title"`    // ฟิลด์ Title สำหรับชื่อสินค้าที่ผู้ใช้ต้องการส่งมา
	Price    float64 `json:"price"`    // ฟิลด์ Price สำหรับราคาของสินค้าที่ผู้ใช้กำหนด
	Quantity int     `json:"quantity"` // ฟิลด์ Quantity สำหรับจำนวนสินค้าที่ผู้ใช้กำหนด
}

// โครงสร้างข้อมูลที่ใช้เมื่อรับคำขอในการอัปเดตสถานะของสินค้า
type RequestUpdateItemStatus struct {
	Status constant.ItemStatus `json:"status"` // ฟิลด์ Status สำหรับอัปเดตสถานะของสินค้าจากผู้ใช้ ใช้ enum จาก constant
}
