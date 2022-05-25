package graph

import (
	repo "github.com/mmmommm/go-gql/repository/article"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Article repo.ArticleRepository
}
