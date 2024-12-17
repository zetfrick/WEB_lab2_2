package repository

import (
	"errors"
	"news-service/models"
	"sync"
)

// ErrNewsNotFound ошибка, если новость не найдена
var ErrNewsNotFound = errors.New("news not found")

// Структура репозитория
type NewsRepository struct {
	newsList  []models.News // Список
	idCounter int           // Счетчик
	mutex     *sync.Mutex   // Мьютекс
}

// Создание нового экземпляря NewsRepository
func NewNewsRepository() *NewsRepository {
	return &NewsRepository{
		newsList:  make([]models.News, 0, 50), 
		idCounter: 1,                          
		mutex:     &sync.Mutex{},              
	}
}

// Все новости из репозитория.
func (r *NewsRepository) GetAll() []models.News {
	r.mutex.Lock()          
	defer r.mutex.Unlock()  
	return r.newsList       
}

// Добавляет новость в репозиторий
func (r *NewsRepository) Create(news models.News) models.News {
	r.mutex.Lock()  
	defer r.mutex.Unlock()
	news.ID = r.idCounter
	r.idCounter++
	r.newsList = append(r.newsList, news)
	return news
}

// Обновляет новость в репозитории
func (r *NewsRepository) Update(id int, news models.News) (models.News, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for i, item := range r.newsList {
		if item.ID == id {
			r.newsList[i] = news
			return news, nil
		}
	}
	return models.News{}, ErrNewsNotFound
}

// Удаляет новость из репозитория по ID
func (r *NewsRepository) Delete(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for i, item := range r.newsList {
		if item.ID == id {
			r.newsList = append(r.newsList[:i], r.newsList[i+1:]...)
			return nil
		}
	}
	return ErrNewsNotFound
}
