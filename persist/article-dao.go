package persist

import (
	"ManagementSystem/persist/models"
	sm "ManagementSystem/services/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ArticleDao struct {
	db *gorm.DB
}

func NewArticleDAO() *ArticleDao {
	return &ArticleDao{db: GetDB()}
}

func (dao *ArticleDao) CreateArticle(article *sm.Article) error {
	return dao.db.Create(dao.toArticleDbo(article)).Error
}

func (dao *ArticleDao) GetArticle(id uuid.UUID) (*sm.Article, error) {
	var article models.ArticleDbo
	if err := dao.db.First(&article, id).Error; err != nil {
		return nil, err
	}
	return dao.toArticle(&article), nil
}

func (dao *ArticleDao) GetAllArticles() ([]sm.Article, error) {
	var articleDbos []models.ArticleDbo
	if err := dao.db.Find(&articleDbos).Error; err != nil {
		return nil, err
	}

	var articles []sm.Article
	for _, articleDbo := range articleDbos {
		articles = append(articles, *dao.toArticle(&articleDbo))
	}

	return articles, nil
}

func (dao *ArticleDao) UpdateArticle(article *sm.Article) error {
	return dao.db.Save(dao.toArticleDbo(article)).Error
}

func (dao *ArticleDao) DeleteArticle(id uuid.UUID) error {
	return dao.db.Delete(&models.ArticleDbo{}, id).Error
}

func (dao *ArticleDao) toArticleDbo(article *sm.Article) *models.ArticleDbo {
	return &models.ArticleDbo{
		ID:          article.Id,
		Title:       article.Title,
		Description: article.Description,
		CreatedAt:   article.CreatedAt,
		UserID:      article.UserId,
	}
}

func (dao *ArticleDao) toArticle(article *models.ArticleDbo) *sm.Article {
	return &sm.Article{
		Id:          article.ID,
		Title:       article.Title,
		Description: article.Description,
		CreatedAt:   article.CreatedAt,
		UserId:      article.UserID,
	}
}
