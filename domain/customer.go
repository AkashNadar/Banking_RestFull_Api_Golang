package domain

import "github.com/banking/errors"

type Customer struct {
	Id      string `db:"customer_id"`
	Name    string
	City    string
	Pincode string
	Dob     string `db:"date_of_birth"`
	Status  string
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errors.AppError)
	ById(string) (*Customer, *errors.AppError)
}
