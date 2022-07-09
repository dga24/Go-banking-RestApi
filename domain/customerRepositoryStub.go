package domain

import "banking/errs"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) findAll() ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "David", City: "Barcelona", Zipcode: "08025", DateOfBirth: "1994-08.-24", Status: "1"},
		{Id: "1002", Name: "Viktoriia", City: "Odessa", Zipcode: "08020", DateOfBirth: "1996-0-05", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}
