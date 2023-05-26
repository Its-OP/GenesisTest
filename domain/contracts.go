package domain

import "time"

type ICoinService interface {
	GetCurrentRate(currency string, coin string) *Price
	subscribeUser(email string)
	sendEmails()
}

type ICoinClient interface {
	GetRate(currency string, coin string) (float64, time.Time)
}

type IEmailRepository interface {
	AddEmail(email string)
	Save()
}
