package cryptostructs

import (
	"fmt"
	"time"
)

type CandlestickData struct {
	Time   int64   `json:"t"`
	Open   float64 `json:"o"`
	High   float64 `json:"h"`
	Low    float64 `json:"l"`
	Close  float64 `json:"c"`
	Volume float64 `json:"v"`
}

type Candlestick struct {
	Candlestick    []CandlestickData `json:"data"`
	InstrumentName string            `json:"instrument_name"`
	Interval       string            `json:"interval"`
}

type CandlestickResult struct {
	CandlestickResult Candlestick `json:"result"`
}

func (c CandlestickResult) TextOutput() string {
	var p string
	p += fmt.Sprintf("Instrument Name: %s\nInterval %s\n", c.CandlestickResult.InstrumentName, c.CandlestickResult.Interval)
	for _, Candlestick := range c.CandlestickResult.Candlestick {
		p += fmt.Sprintf("-------------------------------------\nTime: %s\nOpen: %.8f\nHigh %.8f\nLow: %.8f\nClose: %.8f\nVolume: %.0f\n",
			time.Unix(Candlestick.Time, 0).Format(time.RFC822Z), Candlestick.Open, Candlestick.High, Candlestick.Low, Candlestick.Close, Candlestick.Volume)
	}
	return p
}
