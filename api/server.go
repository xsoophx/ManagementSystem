package api

import (
	"ManagementSystem/persist"
	"ManagementSystem/services"
)

type Server struct {
	UserService    *services.UserService
	ArticleService *services.ArticleService
}

func NewServer(dbUrl string) *Server {
	persist.InitDB(dbUrl)

	userDao := persist.NewUserDao()
	articleDao := persist.NewArticleDAO()

	userService := services.NewUserService(userDao)
	articleService := services.NewArticleService(articleDao)

	return &Server{
		UserService:    userService,
		ArticleService: articleService,
	}
}
