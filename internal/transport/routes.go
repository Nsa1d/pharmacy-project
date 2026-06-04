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
	promoService services.PromocodeService,
	orderService services.OrderService,
	paymentService services.PaymentService,
	userService services.UserService,
	reviewService services.ReviewService,
) {
	cartHandler := NewCartHandler(cartService)
	promoHandler := NewPromocodeHandler(promoService)
	orderHandler := NewOrderHandler(orderService)
	paymentHandler := NewPaymentHandler(paymentService)
	userHandler := NewUserHandler(userService)
	medHandler := NewMedicineHandler(medService)
	categoriesHandler := NewCategoryHandler(categoriesService)
	reviewHandler := NewReviewHandler(reviewService)

	cartHandler.RegisterRoutes(router)
	promoHandler.RegisterRoutes(router)
	orderHandler.RegisterRoutes(router)
	paymentHandler.RegisterRoutes(router)
	medHandler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)
	categoriesHandler.RegisterRoutes(router)
	reviewHandler.RegisterRoutes(router)

}