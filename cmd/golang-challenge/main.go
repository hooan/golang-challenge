package main

import (
	"golang-challenge/internal/transactions"
	mail "golang-challenge/pkg/services"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

    txs := transactions.GetTransactions()
	mail.SendEmail("Transaction Summary", txs.Format())
}