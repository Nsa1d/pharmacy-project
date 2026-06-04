package transport

import (
	"net/http"
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/services"
	"strings"

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
		promo.GET("", h.GetAll)
		promo.POST("", h.CreatePromo)
		promo.GET("/:code", h.GetByCode)
		promo.PATCH("/:code", h.Update)
	}
}

func (h *PromocodeHandler) CreatePromo(c *gin.Context) {
	var req models.PromocodeCreate
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

func (h *PromocodeHandler) GetByCode(c *gin.Context) {
	code := c.Query("code")

	promo, err := h.promocode.GetByCode(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, promo)
}

func (h *PromocodeHandler) Update(c *gin.Context) {
	var req models.PromocodeUpdate
	code := c.Param("code")
	promocode := strings.ToUpper(code)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	promo, err := h.promocode.Update(promocode, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, promo)
}
