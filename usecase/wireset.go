package usecase

import (
	"github.com/google/wire"
	"github.com/mmmommm/go-gql/usecase/article"
)

var Set = wire.NewSet(
	article.ProvideArticleGetCase,
)
