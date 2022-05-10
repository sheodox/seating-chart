package migrate

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	dbService "github.com/sheodox/seating-chart/db"
)

func RunMigrations() error {
	db, err := sql.Open("postgres", dbService.ConnectionString())
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///go/src/app/migrate/migrations",
		"postgres", driver)

	if err != nil {
		return err
	}

	err = m.Up()
	// having nothing to migrate returns an error, ignore it!
	if err == migrate.ErrNoChange {
		return nil
	}

	return err
}
