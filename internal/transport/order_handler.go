package transport

import (
	"net/http"
	"pharmacy-project/internal/apperrors"
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service services.OrderService
}

func NewOrderHandler(service services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) RegisterRoutes(r *gin.Engine) {
	orders := r.Group("/users")
	{
		orders.POST("/:id/orders", h.Post)
		orders.GET("/:id/orders", h.GetOrdersByID)
	}

	r.GET("/orders/:id", h.GetFullOrderByID)
	r.PATCH("/orders/:id/status", h.UpdateStatus)
}

func (h *OrderHandler) Post(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID пользователя должен быть числом"})
		return
	}

	var req models.OrderCreate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.service.CreateOrder(uint(id), &req)
	if err != nil {
		errEx := apperrors.Get(err)
		c.JSON(errEx.StatusCode, gin.H{"error": errEx.Msg})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetFullOrderByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}

	order, err := h.service.GetFullOrder(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetOrdersByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}

	orders, err := h.service.GetByUserID(uint(id))
	if err != nil {
		errEx := apperrors.Get(err)
		c.JSON(errEx.StatusCode, gin.H{"error": errEx.Msg})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}

	var req models.OrderStatusUpdate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(uint(id), req.Status); err != nil {
		errEx := apperrors.Get(err)
		c.JSON(errEx.StatusCode, gin.H{"error": errEx.Msg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "статус был успешно обновлен"})
}
