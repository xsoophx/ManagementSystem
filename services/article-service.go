package services

import (
	"ManagementSystem/persist"
	"ManagementSystem/services/models"
	"ManagementSystem/util"
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

func (s *ArticleService) CreateArticle(title string, userId uuid.UUID, description *string) (*models.Article, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("could not generate UUID: %w", err)
	}

	article := &models.Article{Id: id, Title: title, CreatedAt: util.ToPtr(time.Now()), UserId: userId, Description: description}
	if err := s.dao.CreateArticle(article); err != nil {
		return nil, fmt.Errorf("could not create article: %w", err)
	}
	return article, nil
}

func (s *ArticleService) GetArticle(id uuid.UUID) (*models.Article, error) {
	return s.dao.GetArticle(id)
}

func (s *ArticleService) GetAllArticles() ([]models.Article, error) {
	return s.dao.GetAllArticles()
}
