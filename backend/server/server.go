package server

import (
	"backend/domain/repository"
	"backend/pb"
	"backend/server/interceptor"
	"backend/server/service"
	"backend/usecase"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"google.golang.org/grpc"
)

func NewGRPCServer(repos repository.Repositories) *grpc.Server {
	options := make([]grpc.ServerOption, 0)
	options = append(options, grpc.UnaryInterceptor(
		middleware.ChainUnaryServer(
			interceptor.ValidationInterceptor(),
		)),
	)

	server := grpc.NewServer()

	pb.RegisterArticleServiceServer(server, service.NewArticleService(
		usecase.NewArticleUsecase(repos.ArticleRepository),
	))

	return server
}
