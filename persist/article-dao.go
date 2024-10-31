package persist

import (
	"ManagementSystem/persist/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ArticleDao struct {
	db *gorm.DB
}

func NewArticleDAO() *ArticleDao {
	return &ArticleDao{db: GetDB()}
}

func (dao *ArticleDao) CreateArticle(article *models.ArticleDbo) error {
	return dao.db.Create(article).Error
}

func (dao *ArticleDao) GetArticle(id uuid.UUID) (*models.ArticleDbo, error) {
	var article models.ArticleDbo
	if err := dao.db.First(&article, id).Error; err != nil {
		return nil, err
	}
	return &article, nil
}
