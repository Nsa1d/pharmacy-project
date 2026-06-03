package transport

import (
	"net/http"
	"pharmacy-project/internal/apperrors"
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service services.CartService
}

func NewCartHandler(service services.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) RegisterRoutes(r *gin.Engine) {
	cart := r.Group("/users")
	{
		cart.POST("/:id/cart/items", h.AddItem)
		cart.GET("/:id/cart", h.GetAllItems)
		cart.PATCH("/:id/cart/items/:item_id", h.UpdateItem)
		cart.DELETE("/:id/cart/items/:item_id", h.DeleteItem)
	}
}

func (h *CartHandler) AddItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID пользователя должен быть числом"})
		return
	}

	var req models.CartUpsert
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.AddItem(uint(id), &req)
	if err != nil {
		errEx := apperrors.Get(err)
		c.JSON(errEx.StatusCode, gin.H{"error": errEx.Msg})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (h *CartHandler) GetAllItems(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}

	items, err := h.service.GetAllItems(uint(id))
	if err != nil {
		errEx := apperrors.Get(err)
		c.JSON(errEx.StatusCode, gin.H{"error": errEx.Msg})
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *CartHandler) UpdateItem(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID пользователя должен быть числом"})
		return
	}

	itemID, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "item_id должен быть числом"})
		return
	}

	var req models.CartUpdateQuantity

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = h.service.UpdateQuantity(uint(userID), uint(itemID), &req); err != nil {
		errEx := apperrors.Get(err)
		c.JSON(errEx.StatusCode, gin.H{"error": errEx.Msg})
		return
	}

	if req.Quantity == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "товар удален из корзины"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "количество обновлено"})
	}
}

func (h *CartHandler) DeleteItem(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID пользователя должен быть числом"})
		return
	}

	itemID, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "item_id должен быть числом"})
		return
	}

	if err = h.service.DeleteItem(uint(userID), uint(itemID)); err != nil {
		errEx := apperrors.Get(err)
		c.JSON(errEx.StatusCode, gin.H{"error": errEx.Msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "позиция успешно удалена"})
}
