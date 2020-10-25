package cryptoget

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	cstructs "github.com/leflambeur/crypto-client/cryptostructs"
)

func calculateNonce() string {
	epochNanos := time.Now().UnixNano()
	epochMillis := epochNanos / 1000000
	nonceStr := strconv.FormatInt(epochMillis, 10)
	return nonceStr
}

func GetInstruments() (string, error) {
	URL := "https://api.crypto.com/v2/public/get-instruments" //https://exchange-docs.crypto.com/?python#public-get-instruments
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("GET Failure")
	}
	defer resp.Body.Close()
	var iResp cstructs.InstrumentsResult
	if err := json.NewDecoder(resp.Body).Decode(&iResp); err != nil {
		log.Fatalf("Decode Failure %+v", err)
	}
	return iResp.TextOutput(), nil
}
