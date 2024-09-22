package model

import "backend/internal/constant"

type Item struct {
	ID       uint                `json:"id" gorm:"primaryKey"`
	Title    string              `json:"title"`
	Price    float64             `json:"price"`
	Quantity int                 `json:"quantity"`
	Status   constant.ItemStatus `json:"status" gorm:"default:'PENDING'"`
	OwnerID  uint                `json:"owner_id"`
}

type RequestItem struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type RequestUpdateItemStatus struct {
	Status constant.ItemStatus `json:"status"`
}
