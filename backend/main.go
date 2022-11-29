package main

import (
	"backend/constant"
	"backend/infra"
	"backend/infra/repository/article"
	"backend/server"
	"backend/util"

	domainrepo "backend/domain/repository"

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
	}

	router := gin.Default()

	server.NewApiServer(router, repos)

	router.Run()
	logger.Info("server start!")
}

func loadEnv() {
	viper.AutomaticEnv()
	viper.SetDefault(constant.DBHostEnv, "db")
	viper.SetDefault(constant.DBUserEnv, "root")
	viper.SetDefault(constant.DBNameEnv, "dev")
}
