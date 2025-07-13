package main

import (
	"github.com/johnnyFR26/GoOpportunity/config"
	"github.com/johnnyFR26/GoOpportunity/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errf("error to initialize the app: %v", err)
		return
	}
	router.Initialize()
}
