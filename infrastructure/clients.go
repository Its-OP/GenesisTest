package infrastructure

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type BinanceClient struct {
	client  *http.Client
	baseURL *url.URL
}

func NewBinanceClient() *BinanceClient {
	baseUrl := &url.URL{Scheme: "https", Host: "api.binance.com", Path: "/api/v3"}

	binanceClient := &BinanceClient{
		client:  http.DefaultClient,
		baseURL: baseUrl,
	}

	return binanceClient
}

func (binanceClient *BinanceClient) GetRate(currency string, coin string) (float64, time.Time) {
	path := fmt.Sprintf("ticker/price?symbol=%s%s", coin, currency)
	url := binanceClient.baseURL.ResolveReference(&url.URL{Path: path})

	resp, err := binanceClient.client.Get(url.String())
	timestamp := time.Now()

	if err != nil {
		fmt.Println("Err is", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var result PriceResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	price, _ := strconv.ParseFloat(result.price, 10)

	return price, timestamp
}

type PriceResponse struct {
	symbol string
	price  string
}
