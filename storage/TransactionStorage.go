package storage

import "accounting/model"

type Storage model.User

var Db Storage

func (s *Storage) GetTransaction( id string) (*model.Transaction,error) {
	for _, transaction := range s.Transactions {
		if transaction.Id == id{
			return &transaction, nil
		}
	}
	return nil,model.TransactionError{Status: 404, Msg: "Transaction not found"}

}

func(s *Storage) SaveTransaction (transaction model.Transaction) error{

	s.Transactions = append(s.Transactions,transaction)
	return nil
}


func (s *Storage) GetTransactions() []model.Transaction {

	return s.Transactions
}

func (s *Storage) ShowBalance() float64{
	return s.Balance
}