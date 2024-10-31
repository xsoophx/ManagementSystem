package api

import (
	"ManagementSystem/api/generated"
	"ManagementSystem/util"
	"encoding/json"
	"github.com/google/uuid"
	openapitypes "github.com/oapi-codegen/runtime/types"
	"net/http"
	"time"
)

func (s *Server) GetArticles(w http.ResponseWriter, r *http.Request) {
	user := generated.UserDto{Id: uuid.New(), FirstName: "Arbitrary Name", LastName: "Arbitrary Last Name"}

	resp := []generated.ArticleDto{
		{
			CreatedAt:   time.Now(),
			Description: util.ToPtr("Description of Article 1"),
			Id:          uuid.New(),
			Title:       "Article 1",
			UserId:      user.Id,
		}, {
			CreatedAt:   time.Now(),
			Description: util.ToPtr("Description of Article 2"),
			Id:          uuid.New(),
			Title:       "Article 2",
			UserId:      user.Id,
		},
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) GetArticlesId(w http.ResponseWriter, r *http.Request, id openapitypes.UUID) {
	user := generated.UserDto{Id: uuid.New(), FirstName: "Arbitrary Name", LastName: "Arbitrary Last Name"}

	resp := generated.ArticleDto{
		CreatedAt:   time.Now(),
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
	newArticle.CreatedAt = time.Now()

	// TOOO: add repository logic

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(newArticle)
}

func (s *Server) UpdateArticle(w http.ResponseWriter, r *http.Request, id openapitypes.UUID) {
	user := generated.UserDto{Id: uuid.New(), FirstName: "Arbitrary Name", LastName: "Arbitrary Last Name"}

	resp := generated.ArticleDto{
		CreatedAt:   time.Now(),
		Description: util.ToPtr("Description of Article"),
		Id:          id,
		Title:       "Article",
		UserId:      user.Id,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
