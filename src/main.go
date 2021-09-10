package main

import (
	"jarvis/src/core"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app, err := core.CreateFiberApp()
	if err != nil {
		return // create Jarvis Fiber App failed, exit
	}

	// Run and serve
	listenAddr := ":3000"
	if os.Getenv("EXEC_ENV") == "dev" {
		listenAddr = "localhost" + listenAddr
	} // don't turn on outside network access when developing
	_ = app.Listen(":3000")
}
