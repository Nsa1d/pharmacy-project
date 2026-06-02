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

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("", h.Create)
		users.GET("/:id", h.Get)
	}
}

func (h *UserHandler) Create(c *gin.Context) {
	var req models.UserCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.CreateUser(req)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrUserAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		case errors.Is(err, apperrors.ErrInvalidUserInput):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Get(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrUserNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
