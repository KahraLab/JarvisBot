package main

import (
	"jarvis/src/utils"
	"os"

	log "github.com/sirupsen/logrus"
	"jarvis/src/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// setup logger format
	log.SetFormatter(new(utils.LogrusFormatter))

	// create a web server with multiple platform sdk
	app, err := server.CreateFiberApp()
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
