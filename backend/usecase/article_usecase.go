package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"backend/util"
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

	// 更新日をセット
	for _, article := range articles {
		article.SetUpdatedDate()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, articles)
}

func (u *ArticleUsecase) CreateArticle(c *gin.Context) {
	request := &model.RequestArticle{}
	article := &model.Article{}

	request.ArticleId = util.GetUlid()
	err := c.ShouldBindWith(request, binding.JSON)

	// article構造体にセット
	article.SetArticle(request)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	insertErr := u.articleRepository.CreateArticle(article)

	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, insertErr.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
