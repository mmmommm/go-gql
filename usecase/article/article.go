package article

import (
	"context"

	domain "github.com/mmmommm/go-gql/domain/article"
	repo "github.com/mmmommm/go-gql/repository/article"
)

type ArticleGetCase func(ctx context.Context) ([]*domain.Article, error)

type ArticleCreateCaseInput struct {
	ID      string
	Title   string
	Content string
}

func ProvideArticleGetCase(repo repo.ArticleRepository) ArticleGetCase {
	return func(ctx context.Context) ([]*domain.Article, error) {
		return repo.FindAll(ctx)
	}
}

type ArticleCreateCase func(ctx context.Context, in *ArticleCreateCaseInput) error

func ProvideArticleCreateCase(repo repo.ArticleRepository) ArticleCreateCase {
	return func(ctx context.Context, in *ArticleCreateCaseInput) error {
		return repo.Create(ctx, &domain.Article{
			ID:      in.ID,
			Title:   in.Title,
			Content: in.Content,
		})
	}
}
