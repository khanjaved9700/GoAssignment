package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

//  create a dummy data entery in the databse for testing
func TestCreatePerson(t *testing.T) {

	var jsonReq = []byte(`{"id":1,"FirstName":"javed","LastName":"khan","email":"javed@tftus.com"}`)

	req, err := http.NewRequest("POST", "/createperson/", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	r := httptest.NewRecorder()

	control := http.HandlerFunc(CreatePerson)
	control.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}
}

//  test wheather update work correctly or not...
func TestUpdatePersonByID(t *testing.T) {

	var jsonReq = []byte(`{"id":2,"name":"rajat","age":25,"email":"rajat@tftus.com"}`)

	req, err := http.NewRequest("PUT", "/updateperson/{id}", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	r := httptest.NewRecorder()
	control := http.HandlerFunc(UpdatePerson)
	control.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//  here testing getall func
func TestGetAll(t *testing.T) {

	req, err := http.NewRequest("GET", "/getalldata/", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(GetAllPerson)
	control.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}

}

//  testing getbyid func
func TestGetByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/getdata/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(GetPersonById)
	control.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}

}

// testing wheather data will be deleted by id or not
func TestDeletPersonByID(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/deleteperson/2", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(DeletePerson)
	control.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}

}
