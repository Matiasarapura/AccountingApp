package main


import (
	"accounting/controller"
	"accounting/storage"
	"flag"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main(){
	port := flag.String("p", ":8080", "port")
	balance := flag.Float64("b",1000,"account balance")
	flag.Parse()
	storage.Db.Balance = *balance

	r := mux.NewRouter()
	r.HandleFunc("/transaction",controller.GetTransactions).Methods(http.MethodGet)
	r.HandleFunc("/transaction",controller.DoTransaction).Methods(http.MethodPost)
	r.HandleFunc("/transaction/{id}",controller.GetTransaction).Methods(http.MethodGet)
	r.HandleFunc("/balance",controller.GetBalance).Methods(http.MethodGet)
	r.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	loggedRouter := handlers.LoggingHandler(os.Stdout,r)

	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         *port,
		WriteTimeout: 35 * time.Second,
		ReadTimeout:  35 * time.Second,
	}
	log.Printf("Server on port %s", *port )
	log.Fatal(srv.ListenAndServe())

}




