package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"fmt"
	"time"
)


type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)

}

type DefaultAccountService struct{
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError){
	err := req.Validate()
	if err != nil{
		return nil, err
	}
	a := domain.Account{
		AccountId: "",
		CustomerId: req.Customer_id,
		OpeningDate: time.Now().Format("2006-01-02T15:04:05"),
		AccountType: req.AccountType,
		Amount: req.Amount,
		Status: "1",
	}

	newAccount, err := s.repo.Save(a)
	if err!=nil{
		return nil, err
	}
	response := newAccount.ToNewAccountResponseDto()
	return &response, nil
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError){
	fmt.Println(req)
	err := req.Validate()
	if err != nil{
		return nil, err
	}
	if req.IsTransactionTypeWithdrawal(){
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil{
			return nil, err
		}
		if !account.CanWidthdraw(req.Amount){
			return nil, errs.NewValidationError("Insuficient balance in the account")
		}
	}
	t := domain.Transaction{
		AccountId: req.AccountId,
		Amount: req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02T15:04:05"),
	}
	
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil{
		return nil,appError
	}
	response := transaction.ToDto()
	return &response, nil

}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService{
	return DefaultAccountService{repository}
}
