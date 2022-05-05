package echo

import (
	"github.com/mmmommm/go-gql/handler"
	"github.com/labstack/echo/v4"
)

type EchoServer = *echo.Echo

func ProvideEchoServer(h *handler.Handler) EchoServer {
	e := echo.New()
	api := e.Group("/api/v1")
	// user
	e.GET("/article", echo.HandlerFunc(h.ArticleGetHandler))
	api.GET("/health", echo.HandlerFunc(h.HealthHandler))
	return e
}