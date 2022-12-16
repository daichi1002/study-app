package server

import (
	"backend/domain/repository"
	"backend/usecase"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v48/github"
)

func NewApiServer(ctx context.Context, router *gin.Engine, repos repository.Repositories, client *github.Client) {
	articleUsecase := usecase.NewArticleUsecase(repos.ArticleRepository)
	tagUsecase := usecase.NewTagUsecase(repos.TagRepository)
	userUsecase := usecase.NewUserUsecase(repos.UserRepository)

	router.GET("/article/list", func(c *gin.Context) { articleUsecase.GetArticles(c) })
	router.GET("/tags", func(c *gin.Context) { tagUsecase.GetTags(c) })
	router.POST("/article/create", func(c *gin.Context) { articleUsecase.CreateArticle(c) })
	router.GET("/article/:id", func(c *gin.Context) { articleUsecase.GetArticle(c) })
	router.PUT("/article/edit/:id", func(c *gin.Context) { articleUsecase.UpdateArticle(c) })
	router.DELETE("/article/delete/:id", func(c *gin.Context) { articleUsecase.DeleteArticle(c) })
	router.GET("/users", func(c *gin.Context) { userUsecase.GetUsers(c) })
	router.GET("/user/:id", func(c *gin.Context) { userUsecase.GetUser(c, ctx, client) })
	router.GET("/", func(c *gin.Context) { articleUsecase.GetRandomArticle(c) })
	router.POST("/login", func(c *gin.Context) { userUsecase.Login(c) })
}
