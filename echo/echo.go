package echo

import (
	"database/sql"
	"net/http"

	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/mmmommm/go-gql/graph"
	"github.com/mmmommm/go-gql/graph/generated"
	"github.com/mmmommm/go-gql/handler"
	"github.com/mmmommm/go-gql/repository/article"
)

type EchoServer = *echo.Echo

func ProvideEchoServer(h *handler.Handler, db *sql.DB) EchoServer {
	e := echo.New()
	e.GET("/", echo.HandlerFunc(func(c echo.Context) error {
		return c.String(http.StatusOK, "OK\n")
	}))
	api := e.Group("/api/v1")
	graphqlHandler := gqlhandler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{
				Article: article.ProvideArticleRepository(db),
			}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/query")

	api.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	api.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	return e
}
