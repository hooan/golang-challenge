
package repository

import (
 "database/sql"
 "log"

 "golang-challenge/pkg/models"
)

type TxRepository struct {
 DB *sql.DB
}

func NewTxRepository(db *sql.DB)  TxRepositoryInterface{
 return &TxRepository{DB: db}
}

func (m *TxRepository) Insert(post models.Transaction) bool {
 stmt, err := m.DB.Prepare("INSERT INTO transactions (tx_id, tx_date, amount) VALUES ($1, $2, $3)")
 if err != nil {
  log.Println("Insert",err)
  return false
 }
 defer stmt.Close()
 _, err2 := stmt.Exec(post.ID, post.Date, post.Amount)
 if err2 != nil {
  log.Println("Insert",err)
  return false
 }
 return true
}

func (m *TxRepository) Select() []models.Transaction {
 var result []models.Transaction
 rows, err := m.DB.Query("SELECT * FROM transactions")
 if err != nil {
  log.Println("Select",err)
  return nil
 }
 for rows.Next() {
  var (
   tx_id       int
   tx_date    string
   amount    uint16
  )
  err := rows.Scan(&tx_id, &tx_date, &amount)
  if err != nil {
   log.Println(err)
  } else {
   tx := models.Transaction{ID: int(tx_id), Date: tx_date, Amount: float64(amount)}
   result = append(result, tx)
  }
 }
 return result
}