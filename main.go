package main

import (
	"github.com/Enockjoao/Go-Repository/config"
	"github.com/Enockjoao/Go-Repository/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize Configs
	err := config.Init()
	if err == nil {
		logger.Errorf("config Initialization error: %v", err)
		return
	}

	//initialize Ruter
	router.Initialize()
}
