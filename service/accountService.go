package service

import (
	"time"

	"github.com/banking/domain"
	"github.com/banking/dto"
	"github.com/banking/errors"
)

type AccountService interface {
	NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errors.AppError)
	MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s *DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errors.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	acc, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	accDto := acc.ToNewAccountResponseDto()
	return &accDto, nil
}

func (s *DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	if req.IsTransactionTypeWithDrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithDraw(req.Amount) {
			return nil, errors.NewValidationError("Insufficient balance in the account")
		}
	}

	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	transaction, appErr := s.repo.SaveTransaction(t)
	if appErr != nil {
		return nil, appErr
	}

	response := transaction.ToDto()
	return &response, nil

}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{
		repo: repository,
	}
}
