package api

import (
	"ManagementSystem/api/generated"
	"ManagementSystem/util"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (Server) GetArticles(w http.ResponseWriter, r *http.Request) {
	user := generated.UserDto{Id: uuid.New(), Name: "Arbitrary Name"}

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
