package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/olegrom32/tech-assignment/config/wire"
)

// @title Tech assignment API
// @version v1.0.0
func main() {
	_ = godotenv.Load()

	app, err := wire.Initialize()
	if err != nil {
		return
	}

	log.Fatal(app.Run())
}
