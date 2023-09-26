package app

import "github.com/urfave/cli"

// Generate will return the app
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI app"
	app.Usage = "Search IPs and Names of Servers"

	return app
}
