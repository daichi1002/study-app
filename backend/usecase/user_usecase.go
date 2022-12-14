package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
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

func (u *UserUsecase) Login(c *gin.Context) {
	loginInfo := &model.LoginUser{}
	err := c.ShouldBindWith(loginInfo, binding.JSON)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := u.userRepository.GetLoginUser(loginInfo.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// パスワードのチェック
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
