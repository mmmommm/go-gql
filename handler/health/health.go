package health

import (
	"net/http"

	"github.com/mmmommm/go-gql/database/mysql"
	"github.com/labstack/echo/v4"
)

type HealthHandler func(ctx echo.Context) error

func ProvideHealthHandler(mysqlClient mysql.MysqlClient) HealthHandler {
	return func(c echo.Context) error {
		// statusCmd := redisClient.Client.Ping(context.Background())
		// if statusCmd == nil {
		// 	return errors.New("failed to ping to redis server")
		// }
		// if statusCmd.Err() != nil {
		// 	return statusCmd.Err()
		// }

		err := mysqlClient.Ping()
		if err != nil {
			return err
		}

		return c.String(http.StatusOK, "OK\n")
	}
}
