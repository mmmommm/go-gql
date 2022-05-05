package mysql

import (
	"database/sql"
	"errors"

	"github.com/mmmommm/go-gql/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MysqlClient = sql.DB

func ProvideMysqlClient(config *config.Config) (*MysqlClient, error) {
	db, err := sql.Open("mysql", config.DBUser+":"+config.DBPass+"@tcp(localhost:3306)/"+config.DBName)
	if err != nil {
		return nil, err
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
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
	return db, nil
}