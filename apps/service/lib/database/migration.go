package database

import (
	"context"
	"github.com/samsoft00/golang-starter/service/lib/utils"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"go.uber.org/zap"
	"os"
)

func RunMigrationsWithPath(ctx context.Context, db *bun.DB, path string, log *zap.SugaredLogger) {
	log.Info("applying migrations.")

	// discover migrations
	migrations := migrate.NewMigrations()
	err := migrations.Discover(os.DirFS(path))
	if err != nil {
		panic(err)
	}

	// initialize
	migrator := migrate.NewMigrator(db, migrations)
	err = migrator.Init(ctx)
	if err != nil {
		panic(err)
	}

	// migrate
	_, err = migrator.Migrate(ctx)
	if err != nil {
		panic(err)
	}

	log.Info("applied migrations successfully.")
}

// RunMigrations executes migrations from a standard path
func RunMigrations(ctx context.Context, db *bun.DB, log *zap.SugaredLogger) {
	path := utils.GetEnvOr("MIGRATIONS_PATH", "migrations")
	RunMigrationsWithPath(ctx, db, path, log)
}
