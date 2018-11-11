package repository

import (
	"database/sql"
	"fmt"
	"seedapi/model"
	"seedapi/util"
)

var Data DB

type DB interface {
	QueryRow(string, ...interface{}) Row
	Exec(string, ...interface{}) (Result, error)
}

type Row interface {
	Scan(...interface{}) error
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type sqlDB struct {
	db *sql.DB
}

func (s sqlDB) QueryRow(query string, args ...interface{}) Row {
	return s.db.QueryRow(query, args...)
}
func (s sqlDB) Exec(query string, args ...interface{}) (Result, error) {
	return s.db.Exec(query, args...)
}

// SetDatabase set database
func setDatabase(database *sql.DB) {
	Data = &sqlDB{database}
}

func ConnectToDatabase(config *model.AppSetting) *sql.DB {
	db, err := sql.Open(config.DriverName, config.Connection)
	if err != nil {
		util.Logger.Fatal(fmt.Sprintf("Unable to connect to database: %v", err))
		return nil
	}
	setDatabase(db)
	return db
}
