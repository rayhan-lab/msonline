package main

import (
	"office365/config"
	"office365/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Setup()
}

func main() {
	r := routes.InitRoutes()

	if !config.RuntimeConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Run()
}