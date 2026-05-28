package transport

import (
	"errors"
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
		h.handleServiceError(c, err)
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
		switch {
		case errors.Is(err, apperrors.ErrUserNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": apperrors.ErrUserNotFound.Error()})
			return

		case errors.Is(err, apperrors.ErrCartEmpty):
			c.JSON(http.StatusOK, gin.H{"message": err.Error()})
			return

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
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
		h.handleServiceError(c, err)
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
		switch {
		case errors.Is(err, apperrors.ErrUserNotFound),
			errors.Is(err, apperrors.ErrItemNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "позиция успешно удалена"})
}

func (h *CartHandler) handleServiceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, apperrors.ErrUserNotFound),
		errors.Is(err, apperrors.ErrMedicineNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	case errors.Is(err, apperrors.ErrInvalidQuantity),
		errors.Is(err, apperrors.ErrMedicineOutOfStock),
		errors.Is(err, apperrors.ErrInsufficientStock):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	case errors.Is(err, apperrors.ErrItemAlreadyInCart):
		c.JSON(http.StatusOK, gin.H{"error": apperrors.ErrItemAlreadyInCart.Error()})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
