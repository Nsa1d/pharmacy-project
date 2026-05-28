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
	//medRepo := 
	//userRepo := 

	cartService := services.NewCartService(cartRepo, medRepo, userRepo)

	router := gin.Default()
	transport.RegisterRoutes(router, cartService)

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
