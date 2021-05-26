package main

import (
	"otus/app"
)

func main() {
	repository := app.TransactionRepository{}
	parser := app.TransactionParser{}
	//converter := app.CurrencyConverter{}
	converter := app.NewExchangeRatesApiConverter("3a0ecf4091cd8356ae4db601d8168a22")

	budget := app.NewBudgetApplication(repository, parser, converter)
	budget.AddTransaction("Трата -400 RUB Продукты Пятерочка")
	budget.AddTransaction("Трата -1000 RUB Продукты Пятерочка")
	budget.AddTransaction("Прибыль 2000 RUB Продукты Пятерочка")

	budget.OutputTransactions()
	budget.OutputBalanceInCurrency("USD")

}
