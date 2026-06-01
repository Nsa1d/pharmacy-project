package transport

import (
	"pharmacy-project/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	medicineService services.MedicineService,
) {
	medicineHandler := NewMedicineHandler(medicineService)

	medicineHandler.RegisterRoutes(router)
}
