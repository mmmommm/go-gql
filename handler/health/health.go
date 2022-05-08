package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mmmommm/go-gql/database/mysql"
)

type HealthHandler func(ctx echo.Context) error

func ProvideHealthHandler(mysqlClient *mysql.MysqlClient) HealthHandler {
	return func(c echo.Context) error {
		err := mysqlClient.Ping()
		if err != nil {
			return err
		}

		return c.String(http.StatusOK, "OK\n")
	}
}
