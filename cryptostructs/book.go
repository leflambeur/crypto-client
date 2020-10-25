package cryptostructs

import (
        "fmt"
	"time"
)


type BookValues struct {
        Price                   float64
	Quantity		float64
	numberOrders		int

}

type BookData struct {
	BookBids	[]BookValues		`json:"bids"`
	BookAsks	[]BookValues		`json:"asks"`
	BookTime	int64			`json:"t"`
}

type Book struct {
	Book		[]BookData		`json:"data"`
	InstrumentName	string			`json:"instrument_name"`
	BookDepth	int			`json:"depth"`
}

type BookResult struct {
        BookResult      Book			`json:"result"`
}


func (c BookResult) TextOutput() string {
        var p string
	p += fmt.Sprintf("Time: %s\nInstrument Name: %s\n Depth %d\n", time.Unix(c.BookResult.Book.BookTime, 0).Format(time.RFC822Z), c.BookResult.InstrumentName, c.BookResult.BookDepth)
        for i, _ := range c.BookResult.Book.BookBids {
                p += fmt.Sprintf("-------------------------------------\nPrice: %d\nQuantity: %d\nNumber of Orders: %d\n",
                c.BookResult.Book.BookBids[i].Price, c.BookResult.Book.BookBids[i].Quantity, c.BookResult.Book.BookBids[i].numberOrders)
        }
	for j, _ := range c.BookResult.Book.BookAsks {
		p += fmt.Sprintf("-------------------------------------\nPrice: %d\nQuantity: %d\nNumber of Orders: %d\n",
		c.BookResult.Book.BookAsks[j].Price, c.BookResult.Book.BookAsks[j].Quantity, c.BookResult.Book.BookAsks[j].numberOrders)
        }
        return p
}
