package common

import (
	"database/sql"
	"fmt"
)

type MySqlClient struct {
}

func NewSqlClient(source string) *sql.DB {
	db, err := sql.Open("mysql", source)

	if err != nil {
		_ = fmt.Errorf("Failed to open database: %v", err.Error())
		panic(err)
	}

	return db

}
