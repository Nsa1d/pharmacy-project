package pharmacy

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

	if err := db.AutoMigrate(&models.Medicine{}); err != nil {
		log.Fatal("не удалось выполнить миграции: %v", err)
	}

	medicineRepo := repository.NewMedicineRepository(db)

	medicineService := services.NewMedicineService(medicineRepo)

	router := gin.Default()
	transport.RegisterRoutes(router, medicineService)

	if err := router.Run(); err != nil {
		log.Fatal("не удалось запустить HTTP-сервер: %v", err)
	}
}
