package item

import (
	"backend/internal/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateItem(item *model.Item) (uint, error) {
	if err := r.db.Create(item).Error; err != nil {
		return 0, err
	}
	return item.ID, nil
}

func (r *Repository) GetItems() ([]model.Item, error) {
	var items []model.Item
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *Repository) DeleteItem(itemId uint) error {
	// Perform the deletion in the database
	if err := r.db.Delete(&model.Item{}, itemId).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetItemByID(itemId uint) (model.Item, error) {
	var item model.Item
	if err := r.db.First(&item, itemId).Error; err != nil {
		return model.Item{}, err
	}
	return item, nil
}

func (r *Repository) UpdateItem(item model.Item) (model.Item, error) {
	if err := r.db.Save(&item).Error; err != nil {
		return model.Item{}, err
	}
	return item, nil
}

// GetItem ค้นหารายการสินค้าตาม ID

