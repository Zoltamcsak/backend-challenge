package db

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang/glog"
)

func DoMigrations(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		return err
	}

	fsrc, err := (&file.File{}).Open("file://db/migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("file", fsrc, "payroll", driver)

	if err != nil {
		return err
	}

	err = m.Up()
	if err == migrate.ErrNoChange {
		glog.Infof("no migration required")
		return nil
	}

	return err
}
