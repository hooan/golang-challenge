package main

import (
	"context"
	"fmt"
	"golang-challenge/app"
	"golang-challenge/internal/transactions"
	mail "golang-challenge/pkg/services"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

type MyEvent struct {
	Name string `json:"name"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("EXECUTION_ENV") == "AWS" {
		lambda.Start(HandleRequest)
	}

	 var a app.App
		a.CreateConnection()
		a.Migrate()
    txs := transactions.GetTransactions(a.DB)
	mail.SendEmail("Transaction Summary", txs.Format())
	
}

func HandleRequest(ctx context.Context, event *MyEvent) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	var a app.App
    txs := transactions.GetTransactions(a.DB)
	mail.SendEmail("Transaction Summary", txs.Format())
	message := fmt.Sprintf("Hello %s!", event.Name)

	return &message, nil
}