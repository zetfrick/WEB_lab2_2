package news

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
)

// Структура новости
type News struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

// Тут будут новости
var newsList = make([]News, 0, 50)

// ID новостей
var idCounter = 1

// Это надо, чтобы работало хорошо и не ломалось
var mutex = &sync.Mutex{}

// Получить новости
func GetNews(c *gin.Context) {
	c.JSON(200, newsList)
}

// Создать новость
func CreateNews(c *gin.Context) {
	var newNews News
	if err := c.ShouldBindJSON(&newNews); err != nil {
		c.JSON(400, gin.H{"error": err.Error()}) // Ошибка
		return
	}

	mutex.Lock()
	newNews.ID = idCounter // Присваиваем ID
	idCounter++ // Увеличиваем ID
	newsList = append(newsList, newNews) // Добавляем новую новость
	mutex.Unlock()

	c.JSON(201, newNews)
}

// Обновить новости
func UpdateNews(c *gin.Context) {
	idStr := c.Param("id") // Получаем ID новости
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedNews News
	if err := c.ShouldBindJSON(&updatedNews); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	mutex.Lock()
	for i, news := range newsList {
		if news.ID == id {
			newsList[i] = updatedNews // Обновляем новость
			mutex.Unlock()
			c.JSON(200, updatedNews)
			return
		}
	}
	mutex.Unlock()
	c.JSON(404, gin.H{"error": "News not found"})
}

// Удаления новости
func DeleteNews(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	mutex.Lock()
	for i, news := range newsList {
		if news.ID == id {
			newsList = append(newsList[:i], newsList[i+1:]...) // Удаляем новость
			mutex.Unlock()
			c.JSON(200, gin.H{"message": "News deleted"}) // Сообщение об удалении
			return
		}
	}
	mutex.Unlock()
	c.JSON(404, gin.H{"error": "News not found"})
}
