package api

import (
	"ManagementSystem/persist"
	"ManagementSystem/services"
	"context"
)

type Server struct {
	UserService    *services.UserService
	ArticleService *services.ArticleService
	ctx            context.Context
}

func NewServer(dbUrl string, ctx context.Context) *Server {
	persist.InitDB(dbUrl)

	userDao := persist.NewUserDao()
	articleDao := persist.NewArticleDAO()

	userService := services.NewUserService(userDao)
	articleService := services.NewArticleService(articleDao)

	return &Server{
		UserService:    userService,
		ArticleService: articleService,
		ctx:            ctx,
	}
}
