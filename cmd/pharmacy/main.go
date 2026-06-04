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

type noopPurchaseValidator struct{}

func (noopPurchaseValidator) UserPurchasedMedicine(userID, medicineID uint) (bool, error) {
	return true, nil
}

func main() {
	db := config.SetUpDatabaseConnection()

	if err := db.AutoMigrate(
		&models.CartItem{},
		&models.Promocode{},
		&models.User{},
		&models.Medicine{},
		&models.Order{},
		&models.OrderItem{},
		&models.Promocode{},
		&models.Payment{},
		&models.Category{},
		&models.SubCategory{},
	); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	cartRepo := repository.NewCartRepository(db)
	medRepo := repository.NewMedicineRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	userRepo := repository.NewUserRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	promocodeRepo := repository.NewPromocodeRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	reviewRepo := repository.NewReviewRepository(db)

	cartService := services.NewCartService(cartRepo, medRepo, userRepo)
	medService := services.NewMedicineService(medRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	userService := services.NewUserService(userRepo)
	orderService := services.NewOrderService(orderRepo, cartRepo, medRepo, userRepo, promocodeRepo, paymentRepo)
	promocodeService := services.NewPromocodeService(promocodeRepo)
	paymentService := services.NewPaymentService(paymentRepo, orderRepo, cartRepo)
	reviewService := services.NewReviewService(reviewRepo, userRepo, medRepo, noopPurchaseValidator{}, nil)

	router := gin.Default()
	transport.RegisterRoutes(router, medService, categoryService, cartService, promocodeService, orderService, paymentService, userService, reviewService)

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
