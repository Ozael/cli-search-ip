package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the app
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI app"
	app.Usage = "Search IPs and Names of Servers"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search for IPs addresses",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "www.google.com",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}
