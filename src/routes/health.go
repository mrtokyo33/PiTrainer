package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtokyo33/PiTrainer/src/controllers"
)

func InitHealthRoutes(rg *gin.RouterGroup) {
	rg.GET("/health", controllers.HealthCheck)
}
