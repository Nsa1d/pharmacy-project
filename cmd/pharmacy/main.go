package main

import (
	"log"
	"pharmacy-project/internal/config"
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/repository"
	"pharmacy-project/internal/services"
	"pharmacy-project/internal/transport"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()

	db.AutoMigrate(
		&models.CartItem{},
		&models.User{},
		&models.Medicine{},
		&models.Order{},
		&models.OrderItem{},
		&models.Promocode{},
		&models.Payment{},
	)
	cartRepo := repository.NewCartRepository(db)
	medRepo := repository.NewMedicineRepository(db)
	userRepo := repository.NewUserRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	promocodeRepo := repository.NewPromocodeRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)

	cartService := services.NewCartService(cartRepo, medRepo, userRepo)

	orderService := services.NewOrderService(orderRepo, cartRepo, medRepo, userRepo, promocodeRepo, paymentRepo)
	paymentService := services.NewPaymentService(paymentRepo, orderRepo, cartRepo)
	medService := services.NewMedicineService(medRepo)

	router := gin.Default()
	transport.RegisterRoutes(router, medService, cartService, orderService, paymentService)

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
