package postgresql

import (
	"bitbucket.org/ripleyx/import-service/app/shared/infrastructure/log"
	"fmt"
	"os"
)

type Options struct {
	databaseName *string
	host         *string
	port         *int
	user         *string
	password     *string
}

func Config() *Options {
	return new(Options)
}

func (o *Options) DatabaseName(databaseName string) *Options {
	o.databaseName = &databaseName
	return o
}

func (o *Options) Host(host string) *Options {
	o.host = &host
	return o
}

func (o *Options) Port(port int) *Options {
	o.port = &port
	return o
}

func (o *Options) User(user string) *Options {
	o.user = &user
	return o
}

func (o *Options) Password(password string) *Options {
	o.password = &password
	return o
}

func MergeOptions(opts ...*Options) *Options {
	option := new(Options)

	for _, opt := range opts {
		if opt.databaseName != nil {
			option.databaseName = opt.databaseName
		}
		if opt.host != nil {
			option.host = opt.host
		}
		if opt.port != nil {
			option.port = opt.port
		}
		if opt.user != nil {
			option.user = opt.user
		}
		if opt.password != nil {
			option.password = opt.password
		}
	}
	return option
}

var (
	defaultPort = 5432
)

func (o *Options) GetUrlConnection() string {
	UrlCockroachFormat := "postgresql://%v:%v@%v:%v/%v"

	if o.port == nil {
		o.port = &defaultPort
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "local" || environment == "" {
		UrlCockroachFormat = "postgresql://%v:%v@%v:%v/%v?sslmode=disable"
	}

	log.Info("Connection: %s", fmt.Sprintf(UrlCockroachFormat, *o.user, "************", *o.host, *o.port, *o.databaseName))
	return fmt.Sprintf(UrlCockroachFormat, *o.user, *o.password, *o.host, *o.port, *o.databaseName)
}
