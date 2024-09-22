package item

import (
	"backend/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	service Service
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{service: NewService(db)}
}

func (c *Controller) CreateItem(ctx *gin.Context) {
	var request model.RequestItem
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assume userId is retrieved from some auth mechanism
	userId := 1 // Replace with actual user ID

	item, err := c.service.CreateItem(request, uint(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, item)
}

func (c *Controller) GetItems(ctx *gin.Context) {
	items, err := c.service.GetItems()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, items)
}

func (c *Controller) DeleteItem(ctx *gin.Context) {
	// Retrieve the item ID from the URL parameters
	itemIdParam := ctx.Param("id")

	// Convert itemIdParam string to uint
	idUint, err := strconv.ParseUint(itemIdParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID format"})
		return
	}
	itemId := uint(idUint)

	// Call the service to delete the item
	err = c.service.DeleteItem(itemId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message
	ctx.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}

func (c *Controller) PutItem(ctx *gin.Context) {
	// Retrieve the item ID from the URL parameters
	itemIdParam := ctx.Param("id")

	// Convert itemIdParam to uint
	idUint, err := strconv.ParseUint(itemIdParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID format"})
		return
	}
	itemId := uint(idUint)

	// Bind the request JSON to the model.RequestItem struct
	var request model.RequestItem
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to update the item
	updatedItem, err := c.service.UpdateItem(itemId, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the updated item
	ctx.JSON(http.StatusOK, updatedItem)
}

// GetItem ดึงข้อมูลรายการสินค้าตาม ID
func (c *Controller) GetItem(ctx *gin.Context) {
	// Retrieve the item ID from the URL parameters
	itemIdParam := ctx.Param("id")

	// Convert itemIdParam to uint
	idUint, err := strconv.ParseUint(itemIdParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID format"})
		return
	}
	itemId := uint(idUint)

	// Call the service to retrieve the item
	item, err := c.service.GetItem(itemId) // Use the service method
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the item data
	ctx.JSON(http.StatusOK, item)
}