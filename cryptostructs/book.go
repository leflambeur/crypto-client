package cryptostructs

import (
	"fmt"
	"time"
)

type BookData struct {
	BookBids [][]float64 `json:"bids"`
	BookAsks [][]float64 `json:"asks"`
	BookTime int64       `json:"t"`
}

type Book struct {
	Book           []BookData `json:"data"`
	InstrumentName string     `json:"instrument_name"`
	BookDepth      int        `json:"depth"`
}

type BookResult struct {
	BookResult Book `json:"result"`
}

func (c BookResult) TextOutput() string {
	var p string

	for _, book := range c.BookResult.Book {
		p += fmt.Sprintf("Time: %s\nInstrument Name: %s\n Depth %d\n", time.Unix(book.BookTime, 0).Format(time.RFC822Z), c.BookResult.InstrumentName, c.BookResult.BookDepth)
		for _, bid := range book.BookBids {
			if len(bid) == 3 {
				p += fmt.Sprintf("-------------------------------------\nPrice: %.8f\nQuantity: %.2f\nNumber of Orders: %.0f\n",
					bid[0], bid[1], bid[2])
			}
		}
		for _, asks := range book.BookAsks {
			p += fmt.Sprintf("-------------------------------------\nPrice: %.8f\nQuantity: %.2f\nNumber of Orders: %.0f\n",
				asks[0], asks[1], asks[2])
		}
	}
	return p
}
