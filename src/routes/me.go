package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtokyo33/PiTrainer/src/middlewares"
)

func InitMeRoutes(rg *gin.RouterGroup, handler gin.HandlerFunc) {
	me := rg.Group("/me")
	me.Use(middlewares.AuthMiddleware())

	me.GET("/", handler)
}
