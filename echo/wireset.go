package echo

import (
	"github.com/mmmommm/go-gql/handler"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	handler.ProvideHandler,
	ProvideEchoServer,
)