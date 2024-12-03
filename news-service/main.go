package main

import (
	_ "news-service/docs"
	"news-service/news"
	"news-service/repository"
	"news-service/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	// Настройка CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	r.Use(cors.New(config))

	// Настройка Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Инициализация репозитория и сервиса
	newsRepo := repository.NewNewsRepository()
	newsService := service.NewNewsService(newsRepo)
	newsController := news.NewNewsController(newsService)

	// Определение маршрутов
	r.GET("/news", newsController.GetNews)
	r.POST("/news", newsController.CreateNews)
	r.PUT("/news/:id", newsController.UpdateNews)
	r.DELETE("/news/:id", newsController.DeleteNews)

	// Запуск сервера
	r.Run(":8080")
}
