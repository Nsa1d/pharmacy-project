package transport

import (
	"pharmacy-project/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
  medService services.MedicineService,
	cartService services.CartService,
) {
  medHandler := NewMedicineHandler(medicineService)
	cartHadler := NewCartHandler(cartService)

	cartHadler.RegisterRoutes(router)
}
