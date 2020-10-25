package main

import (
	"fmt"
	"log"
	"os"

	cget "github.com/leflambeur/crypto-client/cryptoget"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Commands: []cli.Command{
			{
				Name:    "get-instruments",
				Aliases: []string{"i"},
				Usage:   "Get instruments from crypto.com",
				Action: func(c *cli.Context) error {
					instruments, err := cget.GetInstruments()
					if err != nil {
						log.Println(err)
					}
					fmt.Println(instruments)
					return err
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
