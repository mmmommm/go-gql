//go:build wireinject
// +build wireinject

package go_sample

import (
	"github.com/google/wire"
	"github.com/mmmommm/go-gql/config"
	"github.com/mmmommm/go-gql/database"
	"github.com/mmmommm/go-gql/echo"
	"github.com/mmmommm/go-gql/handler"
	"github.com/mmmommm/go-gql/logger"
	"github.com/mmmommm/go-gql/repository"
	"github.com/mmmommm/go-gql/usecase"
)

type EntryPoint struct {
	Srv    echo.EchoServer
	Config *config.Config
	Logger *logger.Logger
}

func NewEntryPoint() (*EntryPoint, func(), error) {
	wire.Build(
		echo.Set,
		database.Set,
		handler.Set,
		usecase.Set,
		repository.Set,
		config.Set,
		logger.Set,
		wire.Struct(new(EntryPoint), "*"),
	)
	return nil, nil, nil
}
