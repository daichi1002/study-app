package usecase

import (
	"backend/domain/repository"
	"backend/pb"
)

type ArticleUsecase struct {
	articleRepository repository.ArticleRepository
}

func NewArticleUsecase(repo repository.ArticleRepository) *ArticleUsecase {
	return &ArticleUsecase{
		articleRepository: repo,
	}
}

func (u *ArticleUsecase) GetArticles(request *pb.ListArticlesRequest) (*pb.ListArticlesResponse, error) {
	return nil, nil
}

func (u *ArticleUsecase) CreateArticle(request *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	return nil, nil
}
