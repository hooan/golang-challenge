package repository

import 
 "golang-challenge/pkg/models"


type TxRepositoryInterface interface {
 Select() []models.Transaction
 Insert(post models.Transaction) bool
}