package item

import (
	"backend/internal/constant"
	"backend/internal/model"

	"gorm.io/gorm"
)

type Service struct {
	Repository *Repository
}

func NewService(db *gorm.DB) Service {
	return Service{
		Repository: NewRepository(db),
	}
}

func (s Service) CreateItem(req model.RequestItem, userId uint) (model.Item, error) {
	item := model.Item{
		Title:    req.Title,
		Price:    req.Price,
		Quantity: req.Quantity,
		Status:   constant.ItemPendingStatus,
		OwnerID:  userId,
	}

	id, err := s.Repository.CreateItem(&item)
	if err != nil {
		return model.Item{}, err
	}

	item.ID = id
	return item, nil
}

func (s Service) GetItems() ([]model.Item, error) {
	items, err := s.Repository.GetItems()
	if err != nil {
		return nil, err
	}
	return items, nil
}

// Correct DeleteItem service function
func (s Service) DeleteItem(itemId uint) error {
	// Call the repository to delete the item
	err := s.Repository.DeleteItem(itemId)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) UpdateItem(itemId uint, req model.RequestItem) (model.Item, error) {
	// Fetch the existing item
	item, err := s.Repository.GetItemByID(itemId)
	if err != nil {
		return model.Item{}, err
	}

	// Update item fields
	item.Title = req.Title
	item.Price = req.Price
	item.Quantity = req.Quantity

	// Call the repository to save the updated item
	updatedItem, err := s.Repository.UpdateItem(item)
	if err != nil {
		return model.Item{}, err
	}

	return updatedItem, nil
}



// GetItem เรียกใช้ Repository เพื่อดึงข้อมูลรายการสินค้า
func (s Service) GetItem(itemId uint) (model.Item, error) {
	item, err := s.Repository.GetItemByID(itemId)
	if err != nil {
		return model.Item{}, err
	}
	return item, nil
}