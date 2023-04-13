package router

import (
	"salt-cost/internal/controller"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", controller.Ping)

	return r
}
