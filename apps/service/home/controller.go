package home

import (
	"github.com/gin-gonic/gin"
	"github.com/samsoft00/golang-starter/service/lib/logger"
	"go.uber.org/fx"
	"net/http"
)

type ControllerDeps struct {
	fx.In

	Logger *logger.WithLogger
}

type Controller struct {
	*logger.WithLogger
}

func NewController(deps ControllerDeps) *Controller {
	return &Controller{WithLogger: deps.Logger}
}

func (c *Controller) DefaultDisplay(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"user": "samsoft", "secret": "secret"})
}
