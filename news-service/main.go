package main

import (
	_ "news-service/docs"
	"news-service/news"

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
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // HTTP-методы
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	r.Use(cors.New(config)) // Применяем настройки CORS

	// Настройка Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Определение маршрутов
	r.GET("/news", news.GetNews)           // Маршрут для списка новостей
	r.POST("/news", news.CreateNews)       // Маршрут для создания новой новости
	r.PUT("/news/:id", news.UpdateNews)    // Маршрут для обновления новости по ID
	r.DELETE("/news/:id", news.DeleteNews) // Маршрут для удаления новости по ID

	// Запуск сервера
	r.Run(":8080")
}
