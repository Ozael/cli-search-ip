package main

import (
	"cli-search-ip/app"
	"log"
	"os"
)

func main() {
	cliApp := app.Generate()
	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
