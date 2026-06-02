package transport

import (
	"net/http"
	"pharmacy-project/internal/apperrors"
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service services.PaymentService
}

func NewPaymentHandler(service services.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) RegisterRoutes(r *gin.Engine) {
	payments := r.Group("/orders")
	{
		payments.POST("/:id/payments", h.Create)
		payments.GET("/:id/payments", h.GetAll)
	}
	r.GET("/payments/:id", h.GetByID)
}

func (h *PaymentHandler) Create(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}

	var req models.PaymentCreate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment, err := h.service.CreatePayment(uint(id), &req)
	if err != nil {
		errEx := apperrors.Get(err)
		c.JSON(errEx.StatusCode, gin.H{"error": errEx.Msg})
		return
	}

	c.JSON(http.StatusCreated, payment)
}

func (h *PaymentHandler) GetAll(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}

	payments, err := h.service.GetAllPayments(uint(id))
	if err != nil {
		errEx := apperrors.Get(err)
		c.JSON(errEx.StatusCode, gin.H{"error": errEx.Msg})
		return
	}

	c.JSON(http.StatusOK, payments)
}

func (h *PaymentHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}

	payments, err := h.service.GetByID(uint(id))
	if err != nil {
		errEx := apperrors.Get(err)
		c.JSON(errEx.StatusCode, gin.H{"error": errEx})
		return
	}

	c.JSON(http.StatusOK, payments)
}