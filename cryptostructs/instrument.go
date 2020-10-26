package cryptostructs

import (
	"fmt"
)

//https://exchange-docs.crypto.com/?python#public-get-instruments

type InstrumentValues struct {
	Name      string `json:"instrument_name"`
	QCurrency string `json:"quote_currency"`
	BCurrency string `json:"base_currency"`
	Price     int    `json:"price_decimals"`
	Quantity  int    `json:"quantity_decimals"`
}

type Instruments struct {
	Instruments []InstrumentValues `json:"instruments"`
}

type InstrumentsResult struct {
	InstrumentsResult Instruments `json:"result"`
}

func (c InstrumentsResult) TextOutput() string {
	var p string
	for i, _ := range c.InstrumentsResult.Instruments {
		p += fmt.Sprintf("-------------------------------------\nName: %s\nQuote Currency: %s\nBase Currency: %s\nPrice: %d\nQuantity: %d\n",
			c.InstrumentsResult.Instruments[i].Name, c.InstrumentsResult.Instruments[i].QCurrency,
			c.InstrumentsResult.Instruments[i].BCurrency, c.InstrumentsResult.Instruments[i].Price,
			c.InstrumentsResult.Instruments[i].Quantity)
	}
	return p
}
