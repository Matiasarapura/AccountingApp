package model

import "fmt"

type TransactionError struct {
	Status int
	Msg    string
}

func (f TransactionError) Error()string{
	return fmt.Sprintf("An error ocurred: %s" ,f.Msg)
}

