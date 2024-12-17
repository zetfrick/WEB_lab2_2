package service

import (
	"news-service/models"
	"news-service/repository"
)

// Структура управления новостями
type NewsService struct {
	repo *repository.NewsRepository // Репозиторий для доступа к новостям
}

// Создает новый экземпляр NewsService
func NewNewsService(repo *repository.NewsRepository) *NewsService {
	return &NewsService{repo: repo}
}

// Возвращает все новости из репозитория
func (s *NewsService) GetAll() []models.News {
	return s.repo.GetAll() // Получение всех новостей
}

// Добавляет новую новость в репозиторий
func (s *NewsService) Create(news models.News) models.News {
	return s.repo.Create(news)
}

// Обновляет существующую новость в репозитории
func (s *NewsService) Update(id int, news models.News) (models.News, error) {
	return s.repo.Update(id, news)
}

// Удаляет новость из репозитория по ID
func (s *NewsService) Delete(id int) error {
	return s.repo.Delete(id)
}
