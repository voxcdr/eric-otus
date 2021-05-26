package app

import (
	"strconv"
	"strings"
	"time"
)

type ITransactionParser interface {
	Parse(input string) Transaction
}

type TransactionParser struct {
}

func (t TransactionParser) Parse(input string) Transaction {
	//"Трата -400 RUB Продукты Пятерочка"
	//"Трата -2000 RUB Бензин IRBIS"
	//"Трата -500 RUB Кафе Шоколадница"
	res := strings.Split(input, " ")

	amount, _ := strconv.ParseFloat(res[1], 32)
	trx := Transaction{
		Date: time.Now(),
		Amount: CurrencyAmount{
			CurrencyCode: res[2],
			Amount:       float32(amount),
		},
	}
	return trx
}
