package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/mmmommm/go-gql/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MysqlClient = sql.DB

func ProvideMysqlClient(config *config.Config) (*MysqlClient, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=%s",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBName,
		config.DBTimeZone,
	)
	client, err := sql.Open(config.DBEngine, dataSourceName)
	if err != nil {
		return nil, err
	}
	driver, err := mysql.WithInstance(client, &mysql.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./database/migrations",
		config.DBName,
		driver,
	)
	if err != nil {
		return nil, err
	}

	err = m.Up()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return nil, err
		}
	}
	
	log.Printf("database running on %s\n", config.DBAddr())
	return client, nil
}
