package common

import (
	"database/sql"
)

type MySqlClient struct {
	*sql.DB
}

func NewSqlClient(source string) *MySqlClient {
	db, err := sql.Open("mysql", source)

	if err != nil {
		Logger().Errorf("Failed to open data base. Reason: %v", err)
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		Logger().Errorf("Failed to connect to mysql. Reason: %v", err)
	}

	return &MySqlClient{db}
}
