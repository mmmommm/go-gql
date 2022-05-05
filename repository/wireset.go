package repository

import (
	"github.com/google/wire"
	"github.com/mmmommm/go-gql/repository/article"
)

var Set = wire.NewSet(
	article.ProvideArticleRepository,
)