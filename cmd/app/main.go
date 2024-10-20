package main

import (
	"api/config"
	"api/internal/app"
)

func main() {
	config := config.LoadConfig()
	app.Run(config)
}
