package models

import "fmt"

type Transaction struct {
	ID          int   `json:"id"`
	Date string `json:"date"`
	Amount      float64 `json:"amount"`
}

type MonthlyTransaction struct {
	Month string `json:"month"`
	Total float64 `json:"total"`
}
type TotalTransaction struct {
	Total float64 `json:"total"`
	Monthly []MonthlyTransaction `json:"monthly"`
	Credit float64 `json:"credit"`
	Debit float64 `json:"debit"`
}

func (t *TotalTransaction) Format() string {
	var formatted string
	formatted += "<IMG SRC='stori_savvi.svg'> <H1>Transaction Summary</H1>\n"
	formatted += "<table border='0'>\n"
	formatted +="<tr>\n"
	formatted += fmt.Sprintf("<td>Total Balance: %.2f <td>\n", t.Total)
	formatted += "</tr><tr>\n"
	formatted += fmt.Sprintf("<td>Average credit: %.2f </td>\n", t.Credit)
	formatted += fmt.Sprintf("<td>Average debit: %.2f </td>\n", t.Debit)
	formatted += "</tr><tr>\n"
	formatted += "<th>Monthly Transactions</th>\n"
	formatted += "</tr><tr>\n"
	for _, monthly := range t.Monthly {
		formatted += fmt.Sprintf("<td>Number of transactions in %s: %.0f </td>\n", monthly.Month, monthly.Total)
			formatted += "</tr><tr>\n"
	}
	formatted += "</tr></table>\n"
	return formatted
}