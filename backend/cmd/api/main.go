package main

import (
	"log"

	"github.com/ilyaGorkun/lineate-shop/internal/server"
)

func main() {
	app := server.NewApp()

	if err := app.Run("8000"); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
