package application

import "btcRate/domain"

type CoinService struct {
	coinClient domain.ICoinClient
	currency   string
	coint      string
}

func (coinService *CoinService) GetCurrentRate(currency string, coin string) *domain.Price {
	rate, time := coinService.coinClient.GetRate(currency, coin)

	return &domain.Price{
		Amount:    rate,
		Currency:  currency,
		Timestamp: time,
	}
}

func NewCoinService(client domain.ICoinClient, currency string) *CoinService {
	service := &CoinService{coinClient: client, currency: currency}

	return service
}
