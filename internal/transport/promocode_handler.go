package transport

import (
	"net/http"
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/services"

	"github.com/gin-gonic/gin"
)

type PromocodeHandler struct {
	promocode services.PromocodeService
}

func NewPromocodeHandler(promocode services.PromocodeService) *PromocodeHandler {
	return &PromocodeHandler{promocode: promocode}
}

func (h *PromocodeHandler) RegisterRoutes(r *gin.Engine) {
	promo := r.Group("/promocode")
	{
		promo.GET("")
		promo.POST("")
	}
}

func (h *PromocodeHandler) CreatePromo(c *gin.Context) {
	var req models.PromocodeUpsert
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	promo, err := h.promocode.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, promo)
}

func (h *PromocodeHandler) GetAll(c *gin.Context) {
	promo, err := h.promocode.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, promo)
}
