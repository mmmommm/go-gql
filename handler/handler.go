package handler

import (
	"github.com/mmmommm/go-gql/handler/health"
	"github.com/mmmommm/go-gql/handler/article"
)

func ProvideHandler(
	articleGet article.ArticleGetHandler,
	health health.HealthHandler,
) *Handler {
	return &Handler{
		HealthHandler: health,
		ArticleGetHandler: articleGet,
	}
}

type Handler struct {
	ArticleGetHandler article.ArticleGetHandler
	HealthHandler health.HealthHandler
}