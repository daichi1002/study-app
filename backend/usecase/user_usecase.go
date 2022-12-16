package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/go-github/v48/github"
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

func (u *UserUsecase) GetUser(c *gin.Context, ctx context.Context, client *github.Client) {
	id := c.Params.ByName("id")
	user, err := u.userRepository.GetUser(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	opt := &github.RepositoryListOptions{Type: "public"}
	repos, _, err := client.Repositories.List(ctx, "daichi1002", opt)

	fmt.Println(repos)
	if _, ok := err.(*github.RateLimitError); ok {
		log.Println("hit rate limit")
	}

	c.JSON(http.StatusOK, user)
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

	// // JWT
	// claims := jwt.StandardClaims{
	// 	Issuer:    user.Id,                               // stringに型変換
	// 	ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
	// }
	// jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// token, err := jwtToken.SignedString([]byte("secret"))
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, map[string]string{"jwt": token})
	c.JSON(http.StatusOK, nil)
}
