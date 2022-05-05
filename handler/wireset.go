package handler

import (
	"github.com/google/wire"
	"github.com/mmmommm/go-gql/handler/article"
	"github.com/mmmommm/go-gql/handler/health"
)

var Set = wire.NewSet(
	article.ProvideArticleGetHandler,
	health.ProvideHealthHandler,
)
