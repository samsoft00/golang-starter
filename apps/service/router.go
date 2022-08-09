package service

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/samsoft00/golang-starter/service/home"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

type deps struct {
	fx.In

	Logger    *zap.SugaredLogger
	GinEngine *gin.Engine

	// add controllers
	HomeController *home.Controller
}

func SetupRoutes(d deps) {
	ginEngine := d.GinEngine
	logger := d.Logger

	// cors
	ginEngine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Content-Length"},
		ExposeHeaders:    []string{"Content-Type", "Content-Length"},
	}))

	{
		ginEngine.GET("/", d.HomeController.DefaultDisplay)
	}

	go func() {
		_ = http.ListenAndServe(":8080", ginEngine)
		logger.Info("start listening")
	}()
}
