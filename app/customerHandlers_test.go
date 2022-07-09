package app

import (
	"banking/dto"
	"banking/errs"
	"banking/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)


var router *mux.Router
var ch CustomerHandlers
var mockService *service.MockCustomerService

func setup(t *testing.T) func(){
	ctrl := gomock.NewController(t)
	mockService = service.NewMockCustomerService(ctrl)
	ch = CustomerHandlers{mockService}
	router = mux.NewRouter()
	router.HandleFunc("/customers", ch.GetAllCustomers)
	return func() {
		router = nil
		defer ctrl.Finish()
	}

}

func Test_return_customers_status_200(t *testing.T) {
	//Arrange
	teardown := setup(t)
	defer teardown()

	dummyCustomers := []dto.CustomerResponse{
		{Id: "1001", Name: "David", City: "Barcelona", Zipcode: "08025", DateOfBirth: "1994-08.-24", Status: "1"},
		{Id: "1002", Name: "Viktoriia", City: "Odessa", Zipcode: "08020", DateOfBirth: "1996-0-05", Status: "1"},
	}
	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers,nil)


	request, _ := http.NewRequest(http.MethodGet,"/customers", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusOK{
		t.Error("Failed while testing the status")
	}
}

func Test_should_return_status_500_error_message(t* testing.T){
		//Arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockService := service.NewMockCustomerService(ctrl)
		mockService.EXPECT().GetAllCustomer("").Return(nil,errs.NewUnexpectedError("some database error"))
		ch := CustomerHandlers{mockService}
	
		router := mux.NewRouter()
		router.HandleFunc("/customers", ch.GetAllCustomers)
		request, _ := http.NewRequest(http.MethodGet,"/customers", nil)
		
		//Act
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		//Assert
		if recorder.Code != http.StatusInternalServerError{
			t.Error("Failed while testing the status")
		}

}