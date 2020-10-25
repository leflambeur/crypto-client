package main

import(
	"os"
	"fmt"
	"log"
	"github.com/urfave/cli"
	cget "cryptoget"
)

func main(){
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:"get-instruments",
				Aliases: []string{"i"},
				Usage: "Get instruments from crypto.com",
				Action: func(c *cli.Context) error {
					instruments, err := cget.GetInstruments()
					if err != nil{
						log.Println(err)
					}
					fmt.Println(instruments)
					return(err)
				},
			},
			{
				Name:"get-book",
				Aliases: []string{"book"},
				Usage: "Get book info for Instrument from crypto.com",
				Action: func(c *cli.Context) error {
					instrumentName := "CRO_BTC"
					Depth := "10"
					book, err := cget.GetBook(instrumentName, Depth)
                                        if err != nil{
                                                log.Println(err)
                                        }
                                        fmt.Println(book)
                                        return(err)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
