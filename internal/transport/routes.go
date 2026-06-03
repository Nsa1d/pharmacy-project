package transport

import (
	"pharmacy-project/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	medService services.MedicineService,
	cartService services.CartService,
	orderService services.OrderService,
	paymentService services.PaymentService,
	userService services.UserService,
) {
	cartHandler := NewCartHandler(cartService)
	orderHandler := NewOrderHandler(orderService)
	paymentHandler := NewPaymentHandler(paymentService)
	userHandler := NewUserHandler(userService)
	medHandler := NewMedicineHandler(medService)

	cartHandler.RegisterRoutes(router)
	orderHandler.RegisterRoutes(router)
	paymentHandler.RegisterRoutes(router)
	medHandler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)
}
