package domain

import "time"

type ICoinService interface {
	GetCurrentRate(currency string, coin string) *Price
	Subscribe(email string)
	SendRateEmails()
}

type ICoinClient interface {
	GetRate(currency string, coin string) (float64, time.Time)
}

type IEmailRepository interface {
	AddEmail(email string)
	GetAll() []string
	Save()
}

type IEmailClient interface {
	Send(recipients []string, htmlContent string)
}
