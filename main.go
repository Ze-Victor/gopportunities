package main

import (
	"fmt"

	"github.com/Ze-Victor/gopportunities/config"
	"github.com/Ze-Victor/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initilization error: %v", err)
		fmt.Println(err)
		return
	}
	// Initialize Router
	router.Initialize()
}
