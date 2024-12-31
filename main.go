package main

import (
	"log"
	"os"

	"github.com/erikdenisrs97/mini-cli-app-go/app"
)

func main() {
	app := app.Generate()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
