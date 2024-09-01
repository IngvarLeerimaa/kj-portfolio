package database

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./data/database/database.db")
	if err != nil {
		return nil, err
	}
	err = MigrateUp()
	if err != nil {
		if err.Error() != "no change" {
			fmt.Println("Migrate: ", err)
		}
	}
	return db, nil
}

func MigrateUp() error {
	m, err := migrate.New("file://./data/database/migrations/", "sqlite3://./data/database/database.db")
	if err != nil {
		return err
	}
	defer m.Close()

	err = m.Up()
	if err != nil {
		return err
	}
	return nil
}
