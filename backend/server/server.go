package server

import (
	"backend/domain/repository"
	"backend/usecase"

	"github.com/gin-gonic/gin"
)

func NewApiServer(router *gin.Engine, repos repository.Repositories) {
	articleUsecase := usecase.NewArticleUsecase(repos.ArticleRepository)
	tagUsecase := usecase.NewTagUsecase(repos.TagRepository)

	// エンドポイント
	router.GET("/article/list", func(c *gin.Context) { articleUsecase.GetArticles(c) })
	router.GET("/tags", func(c *gin.Context) { tagUsecase.GetTags(c) })
	router.POST("/article/create", func(c *gin.Context) { articleUsecase.CreateArticle(c) })
	router.GET("/article/:id", func(c *gin.Context) { articleUsecase.GetArticle(c) })
	router.PUT("/article/edit/:id", func(c *gin.Context) { articleUsecase.UpdateArticle(c) })
	router.DELETE("/article/delete/:id", func(c *gin.Context) { articleUsecase.DeleteArticle(c) })
}
