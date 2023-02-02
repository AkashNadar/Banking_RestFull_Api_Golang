package dto

import (
	"strings"

	"github.com/banking/errors"
)

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerId      string  `json:"-"`
}

func (tr *TransactionRequest) IsTransactionTypeWithDrawal() bool {
	if strings.ToLower(tr.TransactionType) == "withdraw" {
		return true
	} else {
		return false
	}
}

func (tr *TransactionRequest) Validate() *errors.AppError {
	if strings.ToUpper(tr.TransactionType) != "WITHDRAW" && strings.ToUpper(tr.TransactionType) != "DEPOSIT" {
		return errors.NewValidationError("Transaction type must be \"WITHDRAW\" or \"DEPOSIT\"")
	}
	if tr.Amount < 1 {
		return errors.NewValidationError("Amount must be greater than zero")
	}

	return nil
}
