package main

import (
	"banking/app"
	"banking/logger"
)

func main() {
	logger.Info("starting the application")
	app.Start()
}
