package app

import "fmt"

type IBudgetApplication interface {
	AddTransaction(input string)
	OutputTransactions()
	OutputBalanceInCurrency(currencyCode string)
}

type BudgetApplication struct {
	Repository TransactionRepository
	Parser     ITransactionParser
	Converter  ICurrencyConverter
}

func (b *BudgetApplication) AddTransaction(input string) {
	res := b.Parser.Parse(input)
	b.Repository.AddTransaction(res)
}

func (b *BudgetApplication) OutputTransactions() {
	for _, trx := range b.Repository.GetTransactions() {
		fmt.Println(trx.Date, "amount: ", trx.Amount)
	}
}

func (b *BudgetApplication) OutputBalanceInCurrency(currencyCode string) {
	var res float32
	for _, trx := range b.Repository.GetTransactions() {
		res += b.Converter.ConvertCurrency(trx.Amount, currencyCode).Amount
	}
	fmt.Println("Total:", res)
}

func NewBudgetApplication(repository TransactionRepository, parser ITransactionParser, converter ICurrencyConverter) BudgetApplication {
	res := BudgetApplication{
		Repository: repository,
		Parser:     parser,
		Converter:  converter,
	}
	return res
}
