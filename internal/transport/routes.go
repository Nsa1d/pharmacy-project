package transport

import (
	"pharmacy-project/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	cartService services.CartService,
) {
	cartHadler := NewCartHandler(cartService)

	cartHadler.RegisterRoutes(router)
}
