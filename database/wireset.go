package database

import (
	"github.com/google/wire"
	"github.com/mmmommm/go-gql/database/mysql"
)

var Set = wire.NewSet(
	mysql.ProvideMysqlClient,
)
