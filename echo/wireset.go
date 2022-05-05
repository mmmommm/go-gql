package echo

import (
	"github.com/google/wire"
	"github.com/mmmommm/go-gql/handler"
)

var Set = wire.NewSet(
	handler.ProvideHandler,
	ProvideEchoServer,
)
