package postgresql

import (
	"bitbucket.org/ripleyx/import-service/app/shared/infrastructure/log"
	"os"
	"strconv"
)

func AutoMigrateEntities(connection Connection, models ...interface{}) {
	log.Info("AutoMigrateEntities...")
	migrate := NewMigrate(connection)
	migrate.AutoMigrateAll(
		models,
	)
	log.Info("AutoMigrateEntities... OK")
}

func CreateDbConnection() *DbConnection {
	var port int
	var err error

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	if len(dbHost) == 0 || len(dbPort) == 0 || len(dbDatabase) == 0 ||
		len(dbUsername) == 0 || len(dbPassword) == 0 {
		log.Fatal("invalid connection data error")
	}

	if port, err = strconv.Atoi(dbPort); err != nil {
		log.Fatal("invalid port error")
	}

	connection := NewPostgresqlConnection(Config().
		Host(dbHost).
		Port(port).
		DatabaseName(dbDatabase).
		User(dbUsername).
		Password(dbPassword),
	)
	return connection
}
