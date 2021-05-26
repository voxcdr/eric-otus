package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ICurrencyConverter interface {
	ConvertCurrency(amount CurrencyAmount, currencyCode string) CurrencyAmount
}

type CurrencyConverter struct {
}

func (c CurrencyConverter) ConvertCurrency(amount CurrencyAmount, currencyCode string) CurrencyAmount {
	return CurrencyAmount{
		CurrencyCode: currencyCode,
		Amount:       amount.Amount,
	}
}

type ExchangeRatesApiConverter struct {
	Rates map[string]float32
}

func NewExchangeRatesApiConverter(apiKey string) ExchangeRatesApiConverter {
	url := "http://api.exchangeratesapi.io/v1/latest?access_key=" + apiKey + "&base=EUR"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err:", err.Error())
		return ExchangeRatesApiConverter{}
	}
	defer resp.Body.Close()

	var data JsonData
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("could not serialize:", err.Error())
		return ExchangeRatesApiConverter{}
	}

	return ExchangeRatesApiConverter{
		Rates: data.Rates,
	}
}

type JsonData struct {
	Success bool               `json:"success"`
	Base    string             `json:"base"`
	Rates   map[string]float32 `json:"rates"`
}

func (e ExchangeRatesApiConverter) ConvertCurrency(amount CurrencyAmount, currencyCode string) CurrencyAmount {
	var convertAmount float32
	eurAmount := amount.Amount / e.Rates[amount.CurrencyCode]
	if amount.CurrencyCode == "EUR" {
		convertAmount = eurAmount
	} else {
		convertAmount = eurAmount * e.Rates[currencyCode]
	}

	return CurrencyAmount{
		CurrencyCode: currencyCode,
		Amount:       convertAmount,
	}
}
