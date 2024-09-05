package transactions

import (
	"encoding/csv"
	"fmt"
	"golang-challenge/pkg/models"
	"os"
	"strconv"
	"strings"
)

func GetTransactions() *models.TotalTransaction {
	file, err := os.Open(os.Getenv("TRANSACTIONS_FILE"))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil
	}

	var transactions []models.Transaction

	for _, record := range records {
		if(record[0] == "Id") {
			continue
		}

		ID, _ := strconv.Atoi(record[0])
		Date := record[1]
		Amount, _ := strconv.ParseFloat(record[2], 64)

		transaction := models.Transaction{
			ID:     int(ID),
			Date: Date,
			Amount: Amount,
		}
		transactions = append(transactions, transaction)
	}

	return calculateTotal(transactions)
}

func calculateTotal(transactions []models.Transaction) *models.TotalTransaction {

	monthString := map[string]string{
		"1": "January",
		"2": "February",
		"3": "March",
		"4": "April",
		"5": "May",	
		"6": "June",	
		"7": "July",	
		"8": "August",	
		"9": "September",
		"10": "October",
		"11": "November",	
		"12": "December",
	}

	credit := 0.0
	debit := 0.0
	total := 0.0
	monthlyMap := make(map[string]models.MonthlyTransaction)	

	for _, transaction := range transactions {
		total += transaction.Amount
		month := strings.Split(transaction.Date, "/")[0]
		monthly, ok := monthlyMap[month]
		if ok {	
			monthly.Total += 1
			monthlyMap[month] = monthly
		} else {
			monthlyMap[month] = models.MonthlyTransaction{
				Month: monthString[month],
				Total: 1,
			}
		}
		if transaction.Amount > 0 {
				credit += transaction.Amount
			} else {	
				debit += transaction.Amount
			}
	}

	monthly := make([]models.MonthlyTransaction, 0, len(monthlyMap))
	for _, value := range monthlyMap {
		monthly = append(monthly, value)
	}

	return &models.TotalTransaction{
		Total: total,
		Monthly: monthly,
		Credit: credit/float64(len(monthly)),
		Debit: debit/float64(len(monthly)),
	}
}

type TotalTransactionFormatter struct {
	*models.TotalTransaction
}

