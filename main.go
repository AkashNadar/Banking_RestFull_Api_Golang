package main

import (
	"github.com/banking/app"
	"github.com/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
