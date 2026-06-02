package transport

import (
	"pharmacy-project/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	medService services.MedicineService,
	categoriesService services.CategoryService,
	cartService services.CartService,
	userService services.UserService,
	reviewService services.ReviewService,
) {
	medHandler := NewMedicineHandler(medService)
	categoriesHandler := NewCategoryHandler(categoriesService)
	cartHandler := NewCartHandler(cartService)
	userHandler := NewUserHandler(userService)
	reviewHandler := NewReviewHandler(reviewService)

	medHandler.RegisterRoutes(router)
	categoriesHandler.RegisterRoutes(router)
	cartHandler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)
	reviewHandler.RegisterRoutes(router)

}
