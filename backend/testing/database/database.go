package database

import "database/sql"

func NewFakeDB() (*sql.DB, error) {
	// docker mysql
	return sql.Open("mysql", "dab_user:passw0rd@tcp(db:3306)/dab_db?parseTime=true")
}
