package app

import "time"

type CurrencyAmount struct {
	CurrencyCode string
	Amount       float32
}

type Transaction struct {
	Date   time.Time
	Amount CurrencyAmount
}
