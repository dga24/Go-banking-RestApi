package domain

import (
	"banking/dto"
	"banking/errs"
)

type Account struct{
	AccountId string  `db:"account_id"`  
	CustomerId string `db:"customer_id"` 
	OpeningDate string`db:"opening_date"`
	AccountType string`db:"account_type"`
	Amount float64    `db:"amount"` 
	Status string     `db:"status"` 
}

type AccountRepository interface {
	Save(account Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() (dto.NewAccountResponse){
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

func (a Account) CanWidthdraw(amount float64) bool{
	if a.Amount < amount {
		return false
	} 
	return true
}

func NewAccount(customerId, accountType string, amount float64) Account {
	return Account{
		CustomerId:  customerId,
		OpeningDate: "2006-01-02T15:04:05",
		AccountType: accountType,
		Amount:      amount,
		Status:      "1",
	}
}