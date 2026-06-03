package transport

import (
	"pharmacy-project/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	medService services.MedicineService,
	cartService services.CartService,
	userService services.UserService,
	reviewService services.ReviewService,
) {
	medHandler := NewMedicineHandler(medService)
	cartHadler := NewCartHandler(cartService)
	userHandler := NewUserHandler(userService)
	reviewHndler := NewReviewHandler(reviewService)

	cartHadler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)
	medHandler.RegisterRoutes(router)
	reviewHndler.RegisterRoutes(router)

}
