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
	coin    string
}

func NewBinanceClient() *BinanceClient {
	baseUrl := &url.URL{Scheme: "https", Host: "api.binance.com", Path: "/api/v3"}

	binanceClient := &BinanceClient{
		client:  http.DefaultClient,
		baseURL: baseUrl,
		coin:    "BTC",
	}

	return binanceClient
}

func (binanceClient *BinanceClient) getPrice(currency string) (float64, time.Time) {
	path := fmt.Sprintf("ticker/price?symbol=%d%d", binanceClient.coin, currency)
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

type withHeader struct {
	http.Header
	rt http.RoundTripper
}

func WithHeader(rt http.RoundTripper) withHeader {
	if rt == nil {
		rt = http.DefaultTransport
	}

	return withHeader{Header: make(http.Header), rt: rt}
}

func (h withHeader) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(h.Header) == 0 {
		return h.rt.RoundTrip(req)
	}

	req = req.Clone(req.Context())
	for k, v := range h.Header {
		req.Header[k] = v
	}

	return h.rt.RoundTrip(req)
}

type PriceResponse struct {
	symbol string
	price  string
}
