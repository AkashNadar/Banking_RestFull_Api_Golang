package domain

import (
	"github.com/banking/dto"
	"github.com/banking/errors"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

type AccountRepository interface {
	Save(Account) (*Account, *errors.AppError)
	SaveTransaction(Transaction) (*Transaction, *errors.AppError)
	FindBy(accountId string) (*Account, *errors.AppError)
}

func (a Account) CanWithDraw(amount float64) bool {
	return a.Amount < amount
}
