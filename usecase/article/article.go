package article

import (
	"context"

	domain "github.com/mmmommm/go-gql/domain/article"
	repo "github.com/mmmommm/go-gql/repository/article"
)

// type GetCaseInput struct {}
type ArticleGetCase func(ctx context.Context) ([]*domain.Article, error)

func ProvideArticleGetCase(repo repo.ArticleRepository) ArticleGetCase {
	return func(ctx context.Context) ([]*domain.Article, error) {
		return repo.FindAll(ctx)
	}
}