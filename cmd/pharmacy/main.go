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

	if err := db.AutoMigrate(&models.Cart{}, &models.User{}, &models.Medicine{}, &models.Category{}, &models.SubCategory{}); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	cartRepo := repository.NewCartRepository(db)
	medRepo := repository.NewMedicineRepository(db)
	catRepo := repository.NewCategoryRepository(db)
	userRepo := repository.NewUserRepository(db)
	reviewRepo := repository.NewReviewRepository(db)

	cartService := services.NewCartService(cartRepo, medRepo, userRepo)
	medService := services.NewMedicineService(medRepo)
	catService := services.NewCategoryService(catRepo)
	reviewService := services.NewReviewService(reviewRepo, userRepo, medRepo, noopPurchaseValidator{}, nil)
	userService := services.NewUserService(userRepo)

	router := gin.Default()
	transport.RegisterRoutes(router, medService, catService, cartService, userService, reviewService)

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
