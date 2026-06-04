package transport

import (
	"net/http"
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MedicineHandler struct {
	service services.MedicineService
}

func NewMedicineHandler(service services.MedicineService) *MedicineHandler {
	return &MedicineHandler{service: service}
}
func (h *MedicineHandler) RegisterRoutes(r *gin.Engine) {
	medicine := r.Group("/medicine")
	{
		medicine.GET("", h.GetAll)
		medicine.GET("/:id", h.GetByID)
		medicine.POST("", h.CreateMedicine)
		medicine.PATCH("/:id", h.UpdateMedicine)
		medicine.DELETE("/:id", h.DeleteMedicine)
	}
}

func (h *MedicineHandler) GetAll(c *gin.Context) {
	medicine, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medicine)
}

func (h *MedicineHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medicine, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medicine)
}

func (h *MedicineHandler) CreateMedicine(c *gin.Context) {
	var req models.MedCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medicine, err := h.service.CreateMed(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"name":                  medicine.Name,
		"description":           medicine.Description,
		"price":                 medicine.Price,
		"in_stock":              medicine.InStock,
		"stock_quantity":        medicine.StockQuantity,
		"category_id":           medicine.CategoryID,
		"subcategory_id":        medicine.SubcategoryID,
		"manufacturer":          medicine.Manufacturer,
		"prescription_required": medicine.PrescriptionRequired,
	})
}

func (h *MedicineHandler) UpdateMedicine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var req models.MedUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	medicine, err := h.service.PatchMed(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"price":          medicine.Price,
		"in_stock":       medicine.InStock,
		"stock_quantity": medicine.StockQuantity,
	})
}

func (h *MedicineHandler) DeleteMedicine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.service.DeleteByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "лекарство удалено из каталога"})
}