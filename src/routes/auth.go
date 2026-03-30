package routes

import "github.com/gin-gonic/gin"

func InitAuthRoutes(rg *gin.RouterGroup, loginHandler gin.HandlerFunc) {
	auth := rg.Group("/auth")
	auth.POST("/login", loginHandler)
}
