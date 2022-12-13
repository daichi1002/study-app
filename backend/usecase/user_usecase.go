package usecase

import (
	"backend/domain/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: repo,
	}
}

func (u *UserUsecase) GetUsers(c *gin.Context) {
	users, err := u.userRepository.ListUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}
