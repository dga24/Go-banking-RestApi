package app

import (
	"banking/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zipCode" xml:"ZipCode"`
	//Costumer_id string `json:"costumer_id" xml:"costumer_id"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	
	customers, err := ch.service.GetAllCustomer(status)
	if err!= nil{
		writeResponse(w,err.Code,err.AsMessage())
	}else{
		writeResponse(w,http.StatusOK,customers)
	}
	
}

func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	c, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w,err.Code,err.AsMessage())
	} else {
		writeResponse(w,http.StatusOK,c)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}){
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err!=nil{
		panic(err)
	}
}
