package application

import (
	"btcRate/domain"
	"fmt"
)

type CoinService struct {
	coinClient      domain.ICoinClient
	emailClient     domain.IEmailClient
	emailRepository domain.IEmailRepository
	currency        string
	coin            string
}

func NewCoinService(client domain.ICoinClient, emailClient domain.IEmailClient, emailRepository domain.IEmailRepository, currency string) *CoinService {
	service := &CoinService{coinClient: client, emailClient: emailClient, emailRepository: emailRepository, currency: currency}

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

func (coinService *CoinService) SendEmails() {
	emails := coinService.emailRepository.GetAll()

	currentPrice := coinService.GetCurrentRate(coinService.currency, coinService.coin)
	htmlTemplate := `<p><strong>Amount:</strong> %f</p>
	<p><strong>Currency:</strong> %s<p>
	<p><strong>Timestamp:</strong> %s<p>`
	htmlBody := fmt.Sprintf(htmlTemplate, currentPrice.Amount, currentPrice.Currency, currentPrice.Timestamp)

	coinService.emailClient.Send(emails, htmlBody)
}
