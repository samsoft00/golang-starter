package home

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

type ControllerDeps struct {
	fx.In

	Logger *zap.Logger
}

type Controller struct {
	logger *zap.Logger
}

func NewController(deps ControllerDeps) *Controller {
	return &Controller{logger: deps.Logger}
}

func (c *Controller) DefaultDisplay(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"user": "samsoft", "secret": "secret"})
}
