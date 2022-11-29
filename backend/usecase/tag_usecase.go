package usecase

import (
	"backend/domain/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TagUsecase struct {
	tagRepository repository.TagRepository
}

func NewTagUsecase(repo repository.TagRepository) *TagUsecase {
	return &TagUsecase{
		tagRepository: repo,
	}
}

func (u *TagUsecase) GetTags(c *gin.Context) {
	tags, err := u.tagRepository.ListTags()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tags)
}
