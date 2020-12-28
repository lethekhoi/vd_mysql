package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (db *sql.DB, err error) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "123456"
	dbName := "demo6"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		return nil, err
	}

	return
}
