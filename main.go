package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtokyo33/PiTrainer/src/routes"
)

func main() {
	r := gin.Default()

	routes.InitRoutes(r)

	r.Run(":8080")
}
