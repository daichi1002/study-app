package server

import (
	"backend/domain/repository"
	"backend/usecase"

	"github.com/gin-gonic/gin"
)

// TODO
// func NewGRPCServer(repos repository.Repositories) *grpc.Server {
// 	options := make([]grpc.ServerOption, 0)
// 	options = append(options, grpc.UnaryInterceptor(
// 		middleware.ChainUnaryServer(
// 			interceptor.ValidationInterceptor(),
// 		)),
// 	)

// 	server := grpc.NewServer(options...)

// 	pb.RegisterArticleServiceServer(server, service.NewArticleService(
// 		usecase.NewArticleUsecase(repos.ArticleRepository),
// 	))

// 	return server
// }

func NewApiServer(router *gin.Engine, repos repository.Repositories) {
	articleUsecase := usecase.NewArticleUsecase(repos.ArticleRepository)

	// エンドポイント
	router.GET("/list", func(c *gin.Context) { articleUsecase.GetArticles(c) })
	router.POST("/create", func(c *gin.Context) { articleUsecase.CreateArticle(c) })
}
