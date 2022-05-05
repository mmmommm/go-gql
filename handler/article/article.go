package article

import (
	"net/http"

	"github.com/labstack/echo/v4"
	domain "github.com/mmmommm/go-gql/domain/article"
	usecase "github.com/mmmommm/go-gql/usecase/article"
)

type ArticleGetHandler func(ctx echo.Context) error

type ArticleGetParams struct{}

// func (p *ArticleGetParams) toInput() *usecase.ArticleGetCaseInput {
// 	// TODO:
// 	return &usecase.ArticleGetCaseInput{}
// }

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
		// req paramsの解析
		// params := &UserGetParams{}
		// if err := ctx.Bind(params); err != nil {
		// 	// TODO: logging
		// 	return err
		// }
		// usecaseからどんなデータ返す
		result, err := uc(ctx.Request().Context())
		if err != nil {
			return err
		}

		// usecaseを呼ぶ
		return ctx.JSON(http.StatusOK, applyArticleGetView(result))
	}
}
