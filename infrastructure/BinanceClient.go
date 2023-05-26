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
	path := fmt.Sprintf("/ticker/price?symbol=%s%s", coin, currency)
	url := binanceClient.baseURL.String() + path

	resp, _ := binanceClient.client.Get(url)
	timestamp := time.Now()

	respBody, _ := io.ReadAll(resp.Body)

	// TODO add error handling

	var result PriceResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	price, _ := strconv.ParseFloat(result.Price, 10)

	return price, timestamp
}

type PriceResponse struct {
	Symbol string
	Price  string
}
