package server

import (
	"backend/domain/repository"
	"backend/usecase"

	"github.com/gin-gonic/gin"
)

func NewApiServer(router *gin.Engine, repos repository.Repositories) {
	articleUsecase := usecase.NewArticleUsecase(repos.ArticleRepository)

	// エンドポイント
	router.GET("/list", func(c *gin.Context) { articleUsecase.GetArticles(c) })
	router.POST("/create", func(c *gin.Context) { articleUsecase.CreateArticle(c) })
}
