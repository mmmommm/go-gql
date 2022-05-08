package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/mmmommm/go-gql/handler"
)

type EchoServer = *echo.Echo

func ProvideEchoServer(h *handler.Handler) EchoServer {
	e := echo.New()
	api := e.Group("/api/v1")
	{
		api.GET("/article", echo.HandlerFunc(h.ArticleGetHandler))
		api.GET("/health", echo.HandlerFunc(h.HealthHandler))
	}
	return e
}
