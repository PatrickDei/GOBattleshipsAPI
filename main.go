package main

import (
	"github.com/PatrickDei/log-lib/logger"
	"go-battleships/app"
)

func main() {
	logger.Info("Launching app")
	app.Start()
}
