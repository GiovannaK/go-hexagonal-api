package main

import (
	"github.com/GiovannaK/go-hexagonal-api.git/adapter/input/routes"
	"github.com/GiovannaK/go-hexagonal-api.git/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	logger.Info("About to start application")
	router := gin.Default()

	routes.InitRoutes(router)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error trying to init application on port 8080", err)
		return
	}
}
