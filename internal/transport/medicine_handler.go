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
		medicine.GET("")
		medicine.GET("/:id")
		medicine.POST("")
		medicine.PATCH("/:id")
		medicine.DELETE("/:id")
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	medicine, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medicine)
}

func (h *MedicineHandler) CreateMedicine(c *gin.Context) {
	var req models.MedCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateMed(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "лекарство добавлено"})
}

func (h *MedicineHandler) UpdateMedicine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var req models.MedUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.PatchMed(uint(id), req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "данные лекарства обновлены"})
}

func (h *MedicineHandler) DeleteMedicine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.service.DeleteByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "лекарство удалено из каталога"})
}
