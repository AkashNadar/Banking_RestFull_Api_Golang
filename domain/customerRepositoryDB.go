package domain

import (
	"database/sql"
	"fmt"

	"github.com/banking/errors"
	"github.com/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (db CustomerRepositoryDB) FindAll(status string) ([]Customer, *errors.AppError) {

	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSQL := "SELECT customer_id, name, city, pincode, date_of_birth, status FROM customers"
		err = db.client.Select(&customers, findAllSQL)
	} else {
		findAllSQL := "SELECT customer_id, name, city, pincode, date_of_birth, status FROM customers where status=?"
		err = db.client.Select(&customers, findAllSQL, status)
	}
	if err != nil {
		logger.Error(fmt.Sprintf("Error while querying customer table: %s", err))
		return nil, errors.NewUnExpectedError("Unexpected Database Error")
	}

	return customers, nil
}

func (db CustomerRepositoryDB) ById(id string) (*Customer, *errors.AppError) {
	getById := "SELECT customer_id, name, city, pincode, date_of_birth, status FROM customers WHERE customer_id = ?"
	var c Customer

	err := db.client.Get(&c, getById, id)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error(fmt.Sprintf("No rows found in db"))
			return nil, errors.NewNotFoundError("Rows not found in the database")
		} else {
			logger.Error(fmt.Sprintf("Error scanning data :%s", err.Error()))
			return nil, errors.NewUnExpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {

	return CustomerRepositoryDB{
		client: dbClient,
	}
}
