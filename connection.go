package main

import (
	"database/sql"
)

func Connection(connString string) *sql.DB {
	conn, err := sql.Open(database_driver, connString)
	Panic(err)

	err = conn.Ping()
	Panic(err)

	return conn
}
