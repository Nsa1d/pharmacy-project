package transport

import (
	"pharmacy-project/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	cartService services.CartService,
	orderService services.OrderService,
	paymentService services.PaymentService,
) {
	cartHandler := NewCartHandler(cartService)
	orderHandler := NewOrderHandler(orderService)
	paymentHandler := NewPaymentHandler(paymentService)

	cartHandler.RegisterRoutes(router)
	orderHandler.RegisterRoutes(router)
	paymentHandler.RegisterRoutes(router)
}
