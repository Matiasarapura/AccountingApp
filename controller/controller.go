package controller

import (
	"accounting/model"
	"accounting/storage"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
)
const CREDIT = "credit"
const DEBIT  = "debit"

func GetTransactions(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	transactions := storage.Db.GetTransactions()
	if len(transactions) == 0 {
		transErr:= model.TransactionError{Status:404,Msg:"No transactions found"}
		log.Println(transErr.Error())
		w.WriteHeader(transErr.Status)
		_ = json.NewEncoder(w).Encode(transErr)
		return
	}
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(transactions)
	return
}

func GetTransaction(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	transaction,err := storage.Db.GetTransaction(id)
	if err != nil {
		log.Println(err.Error())
		_ = json.NewEncoder(w).Encode(err)
		w.WriteHeader(err.(model.TransactionError).Status)
		return
	}
	w.WriteHeader(200)

	_ = json.NewEncoder(w).Encode(transaction)

	return
}


func DoTransaction(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var transactionRequest model.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&transactionRequest)
	if err != nil {
		log.Println(err.Error())

		http.Error(w,err.Error(),400)
		return

	}
	if transactionRequest.Amount <= 0 {

		transErr:= model.TransactionError{Status:400,Msg:"Invalid transaction amount"}
		w.WriteHeader(transErr.Status)

		_ = json.NewEncoder(w).Encode(transErr)
		log.Println(transErr.Error())


		return
	}
	transactionRequest.Type = strings.ToLower(transactionRequest.Type)

	switch transactionRequest.Type {
		case CREDIT :
			storage.Db.Balance = storage.Db.Balance + transactionRequest.Amount
		case DEBIT:
			if storage.Db.Balance - transactionRequest.Amount < 0{
				transErr:= model.TransactionError{Status:400,Msg:"Invalid transaction amount"}
				w.WriteHeader(transErr.Status)

				_ = json.NewEncoder(w).Encode(transErr)
				log.Println(transErr.Error())
				return
			}
			storage.Db.Balance = storage.Db.Balance - transactionRequest.Amount
		default:

			transErr:= model.TransactionError{Status:400,Msg:"Invalid transaction type"}
			w.WriteHeader(transErr.Status)

			log.Println(transErr.Error())
			_ = json.NewEncoder(w).Encode(transErr)
			return
	}

	transaction := model.Transaction{
		Type: transactionRequest.Type,
		Amount: transactionRequest.Amount,
		Date: time.Now(),
		Id: uuid.NewString(),
	}
	err = storage.Db.SaveTransaction(transaction)
	if err != nil {
		http.Error(w,err.Error(),500)
		return
	}
	w.WriteHeader(201)

	_ = json.NewEncoder(w).Encode(transaction)
	return
}


func GetBalance(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	response := make(map[string]interface{})
	response["balance"] = storage.Db.ShowBalance()
	w.WriteHeader(200)

	_ = json.NewEncoder(w).Encode(response)
	return
}