package pharmacy

import (
	"log"
	"pharmacy-project/internal/config"
	"pharmacy-project/transport"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()

	if err := db.AutoMigrate(); err != nil {
		log.Fatal("не удалось выполнить миграции: %v", err)
	}

	router := gin.Default()
	transport.RegisterRoutes()

	if err := router.Run(); err != nil {
		log.Fatal("не удалось запустить HTTP-сервер: %v", err)
	}
}
