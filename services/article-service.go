package services

import (
	"ManagementSystem/persist"
	"ManagementSystem/persist/models"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type ArticleService struct {
	dao *persist.ArticleDao
}

func NewArticleService(dao *persist.ArticleDao) *ArticleService {
	return &ArticleService{dao: dao}
}

func (s *ArticleService) CreateArticle(title string, userID uuid.UUID) (*models.ArticleDbo, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("could not generate UUID: %w", err)
	}

	article := &models.ArticleDbo{ID: id, Title: title, CreatedAt: time.Now(), UserID: userID}
	if err := s.dao.CreateArticle(article); err != nil {
		return nil, fmt.Errorf("could not create article: %w", err)
	}
	return article, nil
}

func (s *ArticleService) GetArticle(id uuid.UUID) (*models.ArticleDbo, error) {
	return s.dao.GetArticle(id)
}
