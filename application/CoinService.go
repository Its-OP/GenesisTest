package application

import "btcRate/domain"

type CoinService struct {
	coinClient      domain.ICoinClient
	emailRepository domain.IEmailRepository
	currency        string
	coin            string
}

func NewCoinService(client domain.ICoinClient, emailRepository domain.IEmailRepository, currency string) *CoinService {
	service := &CoinService{coinClient: client, emailRepository: emailRepository, currency: currency}

	return service
}

func (coinService *CoinService) GetCurrentRate(currency string, coin string) *domain.Price {
	rate, time := coinService.coinClient.GetRate(currency, coin)

	return &domain.Price{
		Amount:    rate,
		Currency:  currency,
		Timestamp: time,
	}
}

func (coinService *CoinService) Subscribe(email string) {
	coinService.emailRepository.AddEmail(email)
	coinService.emailRepository.Save()
}
