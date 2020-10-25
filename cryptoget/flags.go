package cryptoget

import(
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	_ "bytes"
	"strconv"
	_"io"
	_"os"

	cstructs "cryptostructs"
)

func calculateNonce()(string){
	epochNanos := time.Now().UnixNano()
	epochMillis := epochNanos / 1000000
	nonceStr := strconv.FormatInt(epochMillis, 10)
	return nonceStr
}

func GetInstruments()(string, error) {
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

func GetBook(instrumentName string, bookDepth string)(string, error) {
        URL := fmt.Sprintf("https://api.crypto.com/v2/public/get-book?instrument_name=%s&depth=%s", instrumentName, bookDepth) //https://exchange-docs.crypto.com/?python#public-get-book
        resp, err := http.Get(URL)
        if err != nil {
                log.Fatal("GET Failure")
        }
        defer resp.Body.Close()
	//io.Copy(os.Stdout, resp.Body)
        var bResp cstructs.BookResult
        if err := json.NewDecoder(resp.Body).Decode(&bResp); err != nil {
                log.Fatalf("Decode Failure %+v", err)
        }
        return bResp.TextOutput(), nil

}
