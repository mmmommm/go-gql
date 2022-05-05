package handler

import (
	"github.com/mmmommm/go-gql/handler/article"
	"github.com/mmmommm/go-gql/handler/health"
)

func ProvideHandler(
	articleGet article.ArticleGetHandler,
	health health.HealthHandler,
) *Handler {
	return &Handler{
		HealthHandler:     health,
		ArticleGetHandler: articleGet,
	}
}

type Handler struct {
	ArticleGetHandler article.ArticleGetHandler
	HealthHandler     health.HealthHandler
}
