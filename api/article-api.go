package api

import (
	"ManagementSystem/api/generated"
	"ManagementSystem/services/models"
	"ManagementSystem/util"
	"encoding/json"
	"github.com/google/uuid"
	openapitypes "github.com/oapi-codegen/runtime/types"
	"net/http"
	"time"
)

func (s *Server) GetArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := s.ArticleService.GetAllArticles()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var resp []generated.ArticleDto
	for _, article := range articles {
		resp = append(resp, *articleToArticleDto(&article))
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) GetArticlesId(w http.ResponseWriter, r *http.Request, id openapitypes.UUID) {
	user := generated.UserDto{Id: uuid.New(), FirstName: util.ToPtr("Arbitrary Name"), LastName: util.ToPtr("Arbitrary Last Name")}

	resp := generated.ArticleDto{
		CreatedAt:   util.ToPtr(time.Now()),
		Description: util.ToPtr("Description of Article"),
		Id:          id,
		Title:       "Article",
		UserId:      user.Id,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) CreateArticle(w http.ResponseWriter, r *http.Request) {
	userId, err := util.GetUserIDFromContext(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var newArticle generated.ArticleDto
	if err := json.NewDecoder(r.Body).Decode(&newArticle); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	newArticle.UserId = userId
	newArticle.Id = uuid.New()
	newArticle.CreatedAt = util.ToPtr(time.Now())

	article, err := s.ArticleService.CreateArticle(newArticle.Title, newArticle.UserId, newArticle.Description)

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(articleToArticleDto(article))
}

func (s *Server) UpdateArticle(w http.ResponseWriter, r *http.Request, id openapitypes.UUID) {
	user := generated.UserDto{Id: uuid.New(), FirstName: util.ToPtr("Arbitrary Name"), LastName: util.ToPtr("Arbitrary Last Name")}

	resp := generated.ArticleDto{
		CreatedAt:   util.ToPtr(time.Now()),
		Description: util.ToPtr("Description of Article"),
		Id:          id,
		Title:       "Article",
		UserId:      user.Id,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func articleToArticleDto(article *models.Article) *generated.ArticleDto {
	return &generated.ArticleDto{
		Id:          article.Id,
		Title:       article.Title,
		UserId:      article.UserId,
		Description: article.Description,
		CreatedAt:   article.CreatedAt,
	}
}
