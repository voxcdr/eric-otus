package app

type ITransactionRepository interface {
	AddTransaction(transaction Transaction)
	GetTransactions() []Transaction
}

type TransactionRepository struct {
	Transactions []Transaction
}

func (t *TransactionRepository) AddTransaction(transaction Transaction) {
	t.Transactions = append(t.Transactions, transaction)
}

func (t *TransactionRepository) GetTransactions() []Transaction {
	return t.Transactions
}
