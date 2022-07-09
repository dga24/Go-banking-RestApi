package dto

import (
	"net/http"
	"testing"
)

func Test_transaction_type(t *testing.T) {
	//Arrange
	request := TransactionRequest{
		TransactionType: "invalid transaction type",
	}
	//Act
	AppError := request.Validate()
	//Assert
	if AppError.Message != "Transaction type can only be deposit or withdrawal" {
		t.Error("Invalid message while testing trasaction type")
	}
	if AppError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing trasaction type")
	}

}

func Test_return_error_if_amoun_more_than_zero(t *testing.T){
	//ARRANGE
	request := TransactionRequest{ TransactionType: DEPOSIT, Amount: -100}

	//ACT

	AppError := request.Validate()

	if AppError.Message != "Amount cannot be less than zero"{
		t.Error("Invalid message while validating amount")
	}

}