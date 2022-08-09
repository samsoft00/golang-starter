package setup

import (
	"github.com/samsoft00/golang-starter/service"
	"github.com/samsoft00/golang-starter/service/home"
	"github.com/samsoft00/golang-starter/service/lib"
	"go.uber.org/fx"
)

func GetOptions() []fx.Option {
	return lib.SetupApp(lib.Config{
		ServicePrefix: "",
		ServiceName:   "service",
		CusTomFxOptions: []fx.Option{
			fx.Provide(
				home.NewController,
			),
			fx.Invoke(service.SetupRoutes),
		},
	})
}
