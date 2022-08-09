package database

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/samsoft00/golang-starter/service/lib/utils"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"go.uber.org/zap"
)

func NewPostgres(config *utils.DefaultConfig, log *zap.SugaredLogger) *bun.DB {
	var opt *pgx.ConnConfig
	if config.DatabaseURL != "" {
		parsed, err := pgx.ParseConfig(config.DatabaseURL)
		if err != nil {
			log.Fatalf("could not parse connection string: %s", err.Error())
		}
		opt = parsed
	} else {
		opt = &pgx.ConnConfig{
			Config: pgconn.Config{
				Host:     config.DatabaseHost,
				Port:     uint16(config.DatabasePort),
				User:     config.DatabaseUsername,
				Password: config.DatabasePassword,
				Database: config.DatabaseName,
			},
		}
	}

	opt.PreferSimpleProtocol = true

	pgxConfig, err := pgx.ParseConfig(opt.ConnString())
	if err != nil {
		log.Fatalf("could not connect to database: %s", err.Error())
	}

	sqldb := stdlib.OpenDB(*pgxConfig)

	err = sqldb.Ping()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	return bun.NewDB(sqldb, pgdialect.New())
}
