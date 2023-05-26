package domain

import "time"

type ICoinService interface {
	getCurrentPrice(currency string) Price
	subscribeUser(email string)
	sendEmails()
}

type ICoinClient interface {
	GetRate(currency string, coin string) (float64, time.Time)
}
