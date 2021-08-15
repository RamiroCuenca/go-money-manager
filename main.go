package main

import (
	"github.com/RamiroCuenca/go-money-manager/common"
	migrate "github.com/golang-migrate/migrate"
	mMySql "github.com/golang-migrate/migrate/database/mysql"
)

func main() {
	// Initialize ZAP logger
	_ = common.InitLogger()

	// Create sql client
	client := common.NewSqlClient("ramiro:admin@tcp(localhost:3306)/money_manager")

	makeMigrations(client, "money_manager")

}

func makeMigrations(client *common.MySqlClient, dbName string) {
	driver, err := mMySql.WithInstance(client.DB, &mMySql.Config{})

	if err != nil {
		common.Logger().Errorf("Couldnt make migrations. Reason: %v", err)
	}

	m, _ := migrate.NewWithDatabaseInstance(
		"migrations",
		"mysql",
		driver,
	)
	m.Steps(2)
}
