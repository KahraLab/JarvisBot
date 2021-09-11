package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"jarvis/src/core"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app, err := core.CreateFiberApp()
	if err != nil {
		log.Fatalln(err)
		return // create Jarvis Fiber App failed, exit
	}

	// Run and serve
	listenAddr := ":3000"
	if os.Getenv("EXEC_ENV") == "dev" {
		listenAddr = "localhost" + listenAddr
	} // don't turn on outside network access when developing
	_ = app.Listen(":3000")
}
