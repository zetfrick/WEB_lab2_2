package news

import (
	"github.com/gin-gonic/gin"
	"news-service/models"
	"news-service/service"
	"strconv"
)

// Управление
type NewsController struct {
	service *service.NewsService
}

// Создание NewsController
func NewNewsController(service *service.NewsService) *NewsController {
	return &NewsController{service: service}
}

// Получение всех новостей.
func (c *NewsController) GetNews(ctx *gin.Context) {
	newsList := c.service.GetAll() 
	ctx.JSON(200, newsList)
}

// Создание новой новости
func (c *NewsController) CreateNews(ctx *gin.Context) {
	var newNews models.News // Переменная для хранения новости
	if err := ctx.ShouldBindJSON(&newNews); err != nil { 
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if newNews.Title == "" || newNews.Author == "" || newNews.Content == "" { // Проверка заполнености
		ctx.JSON(400, gin.H{"error": "Title, Author, and Content are required"})
		return
	}
	createdNews := c.service.Create(newNews)
	ctx.JSON(201, createdNews)
}

// Обновления новости
func (c *NewsController) UpdateNews(ctx *gin.Context) {
	idStr := ctx.Param("id") // Получаем ID
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedNews models.News // Переменная для хранения обновленной новости
	if err := ctx.ShouldBindJSON(&updatedNews); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if updatedNews.Title == "" || updatedNews.Author == "" || updatedNews.Content == "" { 
		ctx.JSON(400, gin.H{"error": "Title, Author, and Content are required"})
		return
	}

	news, err := c.service.Update(id, updatedNews) // Обновляем новость
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, news)
}

// Удаления новости.
func (c *NewsController) DeleteNews(ctx *gin.Context) {
	idStr := ctx.Param("id") // Получаем ID
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.service.Delete(id); err != nil { // Удаление новости
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "News deleted"})
}
