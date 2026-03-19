package main

import (
	"gin-orm/config"
	"gin-orm/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoutes(router)

	router.Run(":8000")
}
