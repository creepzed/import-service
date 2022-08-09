package postgresql

import (
	"bitbucket.org/ripleyx/import-service/app/shared/infrastructure/log"
)

type Migrate struct {
	connection Connection
}

func NewMigrate(connection Connection) *Migrate {
	return &Migrate{connection: connection}
}

func (m *Migrate) AutoMigrateAll(tables ...interface{}) {
	db, err := m.connection.GetConnection()
	if err != nil {
		log.WithError(err).Fatal(err.Error())
	}
	db = db.AutoMigrate(tables...)
	if db.Error != nil {
		log.WithError(db.Error).Fatal(db.Error.Error())
	}
}
