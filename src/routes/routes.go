package routes

import "github.com/gin-gonic/gin"

func InitRoutes(eng *gin.Engine) {
	api := eng.Group("/api")

	InitHealthRoutes(api)
}
