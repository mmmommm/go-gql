package article

import (
	"net/http"

	"github.com/labstack/echo/v4"
	domain "github.com/mmmommm/go-gql/domain/article"
	"github.com/mmmommm/go-gql/usecase/article"
	usecase "github.com/mmmommm/go-gql/usecase/article"
)

type ArticleGetHandler func(ctx echo.Context) error

type ArticleView struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ArticleGetView struct {
	Article []*domain.Article `json:"article"`
}

func applyArticleGetView(article []*domain.Article) *ArticleGetView {
	return &ArticleGetView{
		Article: article,
	}
}

func ProvideArticleGetHandler(uc usecase.ArticleGetCase) ArticleGetHandler {
	return func(ctx echo.Context) error {
		articles, err := uc(ctx.Request().Context())
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, applyArticleGetView(articles))
	}
}

type ArticleCreateHandler func(ctx echo.Context) error
type ArticleCreateParams struct {
	Title   string
	Content string
}

func toInput(param *ArticleCreateParams) *article.ArticleCreateCaseInput {
	return &article.ArticleCreateCaseInput{
		Title:   param.Title,
		Content: param.Content,
	}
}

func ProvideArticleCreatehandler(uc usecase.ArticleCreateCase) ArticleCreateHandler {
	return func(ctx echo.Context) error {
		var params ArticleCreateParams
		if err := ctx.Bind(&params); err != nil {
			return err
		}
		if err := uc(ctx.Request().Context(), toInput(&params)); err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, `{"message": "ok"}`)
	}
}
