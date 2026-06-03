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

type ReviewHandler struct {
	service services.ReviewService
}

func NewReviewHandler(service services.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

func (h *ReviewHandler) RegisterRoutes(r *gin.Engine) {
	medicines := r.Group("/medicines")
	{
		medicines.GET("/:id/reviews", h.GetByMedicineID)
		medicines.POST("/:id/reviews", h.Create)
	}

	reviews := r.Group("/reviews")
	{
		reviews.PATCH("/:id", h.Update)
		reviews.DELETE("/:id", h.Delete)
	}
}

func (h *ReviewHandler) Create(c *gin.Context) {
	medicineID, ok := parsePositiveID(c.Param("id"))
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id лекарства должен быть положительным числом"})
		return
	}

	var req models.ReviewCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.service.CreateReview(medicineID, req)
	if err != nil {
		h.handleReviewError(c, err)
		return
	}

	c.JSON(http.StatusCreated, review)
}

func (h *ReviewHandler) GetByMedicineID(c *gin.Context) {
	medicineID, ok := parsePositiveID(c.Param("id"))
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id лекарства должен быть положительным числом"})
		return
	}

	reviews, err := h.service.GetReviewsByMedicineID(medicineID)
	if err != nil {
		h.handleReviewError(c, err)
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func (h *ReviewHandler) Update(c *gin.Context) {
	id, ok := parsePositiveID(c.Param("id"))
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id отзыва должен быть положительным числом"})
		return
	}

	var req models.ReviewUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.service.UpdateReview(id, req)
	if err != nil {
		h.handleReviewError(c, err)
		return
	}

	c.JSON(http.StatusOK, review)
}

func (h *ReviewHandler) Delete(c *gin.Context) {
	id, ok := parsePositiveID(c.Param("id"))
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id отзыва должен быть положительным числом"})
		return
	}

	if err := h.service.DeleteReview(id); err != nil {
		h.handleReviewError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "отзыв удален"})
}

func (h *ReviewHandler) handleReviewError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, apperrors.ErrInvalidReviewInput):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, apperrors.ErrReviewNotAllowed):
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	case errors.Is(err, apperrors.ErrReviewNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, apperrors.ErrUserNotFound),
		errors.Is(err, apperrors.ErrMedicineNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func parsePositiveID(raw string) (uint, bool) {
	id, err := strconv.Atoi(raw)
	if err != nil || id <= 0 {
		return 0, false
	}
	return uint(id), true
}
