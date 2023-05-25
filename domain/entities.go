package domain

import "time"

type Price struct {
	amount    float64
	currency  string
	timestamp time.Time
}
