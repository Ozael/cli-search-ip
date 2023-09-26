package main

import (
	"cli-search-ip/app"
	"log"
	"os"
)

func main() {
	cli_app := app.Generate()
	if err := cli_app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
