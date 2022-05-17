package echo

import (
	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/mmmommm/go-gql/database/mysql"
	"github.com/mmmommm/go-gql/graph"
	"github.com/mmmommm/go-gql/graph/generated"
	"github.com/mmmommm/go-gql/handler"
)

type EchoServer = *echo.Echo

func ProvideEchoServer(h *handler.Handler) EchoServer {
	e := echo.New()
	api := e.Group("/api/v1")
	graphqlHandler := gqlhandler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{DB: mysql.NewMysqlClient}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	{
		api.GET("/article", echo.HandlerFunc(h.ArticleGetHandler))
		api.GET("/health", echo.HandlerFunc(h.HealthHandler))
	}
	return e
}
