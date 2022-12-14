package main

import (
	"backend/constant"
	"backend/infra"
	"backend/infra/repository/article"
	"backend/infra/repository/tag"
	"backend/infra/repository/user"
	"backend/server"
	"backend/util"
	"context"
	"time"

	domainrepo "backend/domain/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var logger = util.NewLogger()

func main() {
	// 環境変数読み込み
	loadEnv()
	// DB初期化処理
	gormHandler := infra.NewGormHandler()

	repos := domainrepo.Repositories{
		ArticleRepository: article.NewArticleRepository(gormHandler),
		TagRepository:     tag.NewTagRepository(gormHandler),
		UserRepository:    user.NewUserRepository(gormHandler),
	}

	router := gin.Default()

	setCors(router)

	ctx := context.Background()

	githubClient := server.NewGithubApiClient(ctx)

	server.NewApiServer(ctx, router, repos, githubClient)

	router.Run()
	logger.Info("server start!")
}

func setCors(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func loadEnv() {
	viper.AutomaticEnv()
	viper.SetDefault(constant.DBHostEnv, "db")
	viper.SetDefault(constant.DBUserEnv, "root")
	viper.SetDefault(constant.DBNameEnv, "dev")
}
