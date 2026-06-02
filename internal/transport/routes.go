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
) {
	medHandler := NewMedicineHandler(medService)
	cartHadler := NewCartHandler(cartService)
	userHandler := NewUserHandler(userService)

	cartHadler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)
	medHandler.RegisterRoutes(router)

}
