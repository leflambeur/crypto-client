package main

import (
	"fmt"
	"log"
	"os"

	cget "github.com/leflambeur/crypto-client/cryptoget"
	"github.com/urfave/cli"
)

func main() {
	var app = &cli.App{

		Name:  "crypto-client",
		Usage: "Interact with crypto.com API",

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
			{
				Name: "get-book",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "instrument",
						Value: "CRO_BTC",
						Usage: "Which instrument you wish to get the Order Book for Default: CRO_BTC",
					},
					cli.StringFlag{
						Name:  "depth",
						Value: "10",
						Usage: "Depth of Order Book History",
					},
				},
				Aliases: []string{"book"},
				Usage:   "Get book info for Instrument from crypto.com",
				Action: func(c *cli.Context) error {
					book, err := cget.GetBook(c.String("instrument"), c.String("depth"))
					if err != nil {
						log.Println(err)
					}
					fmt.Println(book)
					return err
				},
			},
			{
				Name: "get-candlestick",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "instrument",
						Value: "CRO_BTC",
						Usage: "Which instrument you wish to get the Order Book for Default: CRO_BTC",
					},
					cli.StringFlag{
						Name:  "timeframe",
						Value: "5m",
						Usage: "Candlestick Timeframe",
					},
				},
				Aliases: []string{"candlestick"},
				Usage:   "Get Candlestick info for Instrument from crypto.com",
				Action: func(c *cli.Context) error {
					candlestick, err := cget.GetCandlestick(c.String("instrument"), c.String("timeframe"))
					if err != nil {
						log.Println(err)
					}
					fmt.Println(candlestick)
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
