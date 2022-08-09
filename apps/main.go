package main

import (
	"github.com/samsoft00/golang-starter/service/setup"
	"go.uber.org/fx"
)

func main() {
	options := setup.GetOptions()
	app := fx.New(options...)
	app.Run()
}
