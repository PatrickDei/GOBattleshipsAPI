package main

import (
	"github.com/PatrickDei/log-lib/logger"
	"go-battleships/app"
)

// SERVER_ADDRESS=localhost SERVER_PORT=8000 DB_USER=root DB_PASSWORD=root DB_ADDRESS=localhost DB_PORT=3306 DB_NAME=Battleships go run main.go
func main() {
	logger.Info("Launching app")
	app.Start()
}
