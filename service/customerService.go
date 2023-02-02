package service

import (
	"github.com/banking/domain"
	"github.com/banking/errors"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, *errors.AppError)
	GetCustomerById(string) (*domain.Customer, *errors.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errors.AppError) {
	if status == "active" {
		return s.repo.FindAll("1")
	} else if status == "inactive" {
		return s.repo.FindAll("0")
	} else {
		return s.repo.FindAll("")
	}

}

func (s DefaultCustomerService) GetCustomerById(id string) (*domain.Customer, *errors.AppError) {
	return s.repo.ById(id)
}

// Helper function
func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repo: repo,
	}
}
