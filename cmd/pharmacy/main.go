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

	if err := db.AutoMigrate(&models.Cart{}, &models.User{}, &models.Medicine{}); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	cartRepo := repository.NewCartRepository(db)
	medRepo := repository.NewMedicineRepository(db)
	userRepo := repository.NewUserRepository(db)

	cartService := services.NewCartService(cartRepo, medRepo, userRepo)
  medService := services.NewMedicineService(medRepo)
  
	router := gin.Default()
	transport.RegisterRoutes(router, cartService, medService)

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
