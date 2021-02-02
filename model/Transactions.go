package model

import "time"

type Transaction struct {
	Date time.Time
	Id string
	Amount float64
	Type string
}

type TransactionRequest struct {
	Amount float64
	Type   string
}
//negative for debit positive credit

