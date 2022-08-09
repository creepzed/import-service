package postgresql

import (
	"bitbucket.org/ripleyx/import-service/app/shared/infrastructure/log"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const LogMode = true

var connection *gorm.DB

type Connection interface {
	GetConnection() (*gorm.DB, error)
	CloseConnection()
}

type DbConnection struct {
	opts *Options
	url  string
}

var (
	ErrConnectionEmpty = errors.New("error creating connection, empty url")
)

func NewPostgresqlConnection(opts ...*Options) *DbConnection {
	databaseOptions := MergeOptions(opts...)
	url := databaseOptions.GetUrlConnection()
	if url == "" {
		log.Fatal(ErrConnectionEmpty.Error())
	}
	return &DbConnection{
		opts: databaseOptions,
		url:  url,
	}
}

func (r *DbConnection) GetConnection() (*gorm.DB, error) {
	var err error
	if connection == nil || !isAlive() {
		log.Info("Trying to connect to DB")
		connection, err = gorm.Open("postgres", r.url)
		if err != nil {
			log.WithError(err).Error("error trying to connect to DB")
			return nil, err
		} else {
			log.Info("Connected to DB")
		}
	}
	connection.LogMode(LogMode)
	connection.SetLogger(log.Logger())
	return connection, nil
}

func (r *DbConnection) CloseConnection() {
	if err := connection.Close(); err != nil {
		log.WithError(err).Error("error trying to close connection")
	} else {
		log.Info("Connection Closed")
	}
}

func isAlive() bool {
	if err := connection.DB().Ping(); err != nil {
		log.WithError(err).Error("error trying to Ping to Db")
		return false
	}
	return true
}
