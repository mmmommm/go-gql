package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/mmmommm/go-gql/domain/article"
	"github.com/mmmommm/go-gql/graph/generated"
	"github.com/mmmommm/go-gql/graph/model"
)

func ProvideGraphHandler() {

}

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.NewArticle) (*model.Article, error) {
	id := uuid.New().String()
	err := r.Resolver.Article.Create(ctx, &domain.Article{
		ID:      id,
		Title:   input.Title,
		Content: input.Content,
	})
	if err != nil {
		return nil, err
	}
	res := &model.Article{
		ID:      id,
		Title:   input.Title,
		Content: input.Content,
	}
	return res, nil
}

func (r *queryResolver) Articles(ctx context.Context) ([]*model.Article, error) {
	articles, err := r.Article.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var res []*model.Article
	for _, article := range articles {
		res = append(res, &model.Article{
			ID:      article.ID,
			Title:   article.Title,
			Content: article.Content,
		})
	}
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
