package domain

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/banking/errors"
	"github.com/banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (db *AccountRepositoryDB) Save(a Account) (*Account, *errors.AppError) {
	insertQuery := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"

	res, err := db.client.Exec(insertQuery, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error(fmt.Sprintf("Error while creating new Account: %s", err))
		return nil, errors.NewUnExpectedError("Unexpected error from the database")
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Error(fmt.Sprintf("Error while retrieving last inserted id: %s", err))
		return nil, errors.NewUnExpectedError("Unexpected error from the database")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDB(clientDB *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{
		client: clientDB,
	}
}

func (db *AccountRepositoryDB) SaveTransaction(t Transaction) (*Transaction, *errors.AppError) {
	// Starting the database transaction block
	tx, err := db.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account : " + err.Error())
		return nil, errors.NewUnExpectedError("Unexpected Database Error")
	}

	// inserting bank account transaction
	result, _ := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	// updating account balance
	if t.IsWithDrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}

	// In case of error Rollback, and changes from both the tables will be reverted.
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errors.NewUnExpectedError("Unexpected Database Error")
	}

	// Commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for back account : " + err.Error())
		return nil, errors.NewUnExpectedError("Unexpected Database Error")
	}

	// getting the last transaction Id from the transaction table
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last transactionID : " + err.Error())
		return nil, errors.NewUnExpectedError("Unexpected Database Error")
	}

	// Getting the latest account information from the accounts table
	account, appErr := db.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)

	// updating the transaction struct with the latest balance
	t.Amount = account.Amount
	return &t, nil
}

func (db *AccountRepositoryDB) FindBy(accountId string) (*Account, *errors.AppError) {
	getById := "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM accounts WHERE account_id = ?"
	var a Account

	err := db.client.Get(&a, getById, accountId)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error(fmt.Sprint("No rows found in db"))
			return nil, errors.NewNotFoundError("Rows not found in the database")
		} else {
			logger.Error(fmt.Sprintf("Error scanning data :%s", err.Error()))
			return nil, errors.NewUnExpectedError("Unexpected database error")
		}
	}

	return &a, nil
}
