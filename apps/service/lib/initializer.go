package lib

import (
	"context"
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samsoft00/golang-starter/service/lib/database"
	"github.com/samsoft00/golang-starter/service/lib/logger"
	"github.com/samsoft00/golang-starter/service/lib/utils"
	"github.com/uptrace/bun"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"time"
)

func init() {
	setupEnv()
}

var startTime time.Time

func setupEnv() {
	if err := godotenv.Load(); err != nil {
		log.Warnf("Error when reading .env file: %s", err.Error())
	}
}

type Config struct {
	ServicePrefix        string
	ServiceName          string
	CustomDatabaseOption fx.Option
	CusTomFxOptions      []fx.Option
}

func SetupApp(appConfig Config) []fx.Option {
	startTime := time.Now()
	defaultConfig := utils.GetConfig(appConfig.ServicePrefix)
	defaultConfig.ServiceName = appConfig.ServiceName
	fxLog := func() fxevent.Logger {
		return &fxevent.ZapLogger{Logger: logger.NewLogger(defaultConfig)}
	}

	var databaseOption fx.Option
	if appConfig.CustomDatabaseOption != nil {
		databaseOption = appConfig.CustomDatabaseOption
	} else {
		databaseOption = defaultDatabaseOption()
	}

	options := []fx.Option{
		fx.Provide(
			func() *utils.DefaultConfig {
				return defaultConfig
			},
		),
		fx.Provide(logger.NewSugaredLogger, logger.NewWithLogger, logger.NewLogger),
		fx.Provide(setupgin),
		databaseOption,
		migrationsHook(appConfig, startTime),
		fx.WithLogger(fxLog),
	}

	options = append(options, appConfig.CusTomFxOptions...)
	return options
}

func setupgin(config *utils.DefaultConfig, logger *zap.Logger) *gin.Engine {
	if config.Environment == utils.Production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	logger.Info("Setting up a gin default instance")

	r := gin.New()

	return r
}

func defaultDatabaseOption() fx.Option {
	return fx.Provide(func(logger *zap.SugaredLogger, config *utils.DefaultConfig) *bun.DB {
		db := database.NewPostgres(config, logger)
		//middleware
		return db
	})
}

func migrationsHook(
	appConfig Config,
	startTime time.Time,
) fx.Option {
	return fx.Invoke(
		func(
			lifecycle fx.Lifecycle,
			db *bun.DB,
			logger *zap.SugaredLogger,
		) {
			lifecycle.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						database.RunMigrations(ctx, db, logger)
						return nil
					},
					OnStop: func(context.Context) error {
						_ = logger.Sync()
						return nil
					},
				},
			)
		},
	)
}
