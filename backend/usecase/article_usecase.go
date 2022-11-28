package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ArticleUsecase struct {
	articleRepository repository.ArticleRepository
}

func NewArticleUsecase(repo repository.ArticleRepository) *ArticleUsecase {
	return &ArticleUsecase{
		articleRepository: repo,
	}
}

func (u *ArticleUsecase) GetArticles(c *gin.Context) {
	articles, err := u.articleRepository.ListArticles()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, articles)
}

func (u *ArticleUsecase) CreateArticle(c *gin.Context) {
	input := &model.Article{}
	err := c.ShouldBindWith(input, binding.JSON)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	insertErr := u.articleRepository.CreateArticle(input)

	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, insertErr.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
