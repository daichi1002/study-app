package service

import (
	"backend/pb"
	"backend/usecase"
	"context"
)

type ArticleService struct {
	ArticleUsecase *usecase.ArticleUsecase
}

func NewArticleService(articleUsecase *usecase.ArticleUsecase) *ArticleService {
	return &ArticleService{
		ArticleUsecase: articleUsecase,
	}
}

func (s *ArticleService) GetArticles(ctx context.Context, request *pb.ListArticlesRequest) (*pb.ListArticlesResponse, error) {
	return s.ArticleUsecase.GetArticles(request)
}

func (s *ArticleService) CreateArticle(ctx context.Context, request *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	return s.ArticleUsecase.CreateArticle(request)
}
