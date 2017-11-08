package main

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"fmt"
	"time"
	"sync"
	"runtime"
)

type QuoteResponse struct {
	Status string
	Name string
	LastPrice float32
	Change float32
	ChangePercent float32
	Timestamp string
	MSDate float32
	MarketCap int
	Volume int
	ChangeYTD float32
	ChangePercentYTD float32
	High float32
	Low float32
	Open float32

}

func main() {
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(4)
	start := time.Now()

	stockSymbols := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}
	wg.Add(len(stockSymbols))

	for _, symbol := range stockSymbols {
		go func(symbol string) {
			//fmt.Println(symbol)
			resp, err := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=%" + symbol)
			if err != nil {
				fmt.Println("Here1")
				panic(err)
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Here2")
				panic(err)
			}

			quote := QuoteResponse{}
			xml.Unmarshal(body, &quote)

			fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
			wg.Done()
		}(symbol)
	}
	wg.Wait()
	elaspsed := time.Since(start)
	fmt.Printf("Elasped time: %s", elaspsed)
}

