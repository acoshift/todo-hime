package repository

import "database/sql"

// Queryer is the interface for *sql.DB and *sql.Tx
type Queryer interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
	Exec(string, ...interface{}) (sql.Result, error)
}
