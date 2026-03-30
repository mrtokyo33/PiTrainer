package routes

import (
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(rg *gin.RouterGroup, createUser gin.HandlerFunc) {
	users := rg.Group("/users")
	users.POST("/", createUser)
}
