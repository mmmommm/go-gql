// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package golang_example

import (
	"github.com/mmmommm/go-gql/config"
	"github.com/mmmommm/go-gql/database/mysql"
	"github.com/mmmommm/go-gql/echo"
	"github.com/mmmommm/go-gql/handler"
	article3 "github.com/mmmommm/go-gql/handler/article"
	"github.com/mmmommm/go-gql/handler/health"
	"github.com/mmmommm/go-gql/logger"
	"github.com/mmmommm/go-gql/repository/article"
	article2 "github.com/mmmommm/go-gql/usecase/article"
)

// Injectors from wire.go:

func NewEntryPoint() (*EntryPoint, func(), error) {
	configConfig, err := config.ProvideConfig()
	if err != nil {
		return nil, nil, err
	}
	db, err := mysql.ProvideMysqlClient(configConfig)
	if err != nil {
		return nil, nil, err
	}
	articleRepository := article.ProvideArticleRepository(db)
	articleGetCase := article2.ProvideArticleGetCase(articleRepository)
	articleGetHandler := article3.ProvideArticleGetHandler(articleGetCase)
	healthHandler := health.ProvideHealthHandler(db)
	handlerHandler := handler.ProvideHandler(articleGetHandler, healthHandler)
	echoEcho := echo.ProvideEchoServer(handlerHandler)
	sugaredLogger, cleanup, err := logger.ProvideZapLogger(configConfig)
	if err != nil {
		return nil, nil, err
	}
	entryPoint := &EntryPoint{
		Srv:    echoEcho,
		Config: configConfig,
		Logger: sugaredLogger,
	}
	return entryPoint, func() {
		cleanup()
	}, nil
}

// wire.go:

type EntryPoint struct {
	Srv    echo.EchoServer
	Config *config.Config
	Logger *logger.Logger
}
