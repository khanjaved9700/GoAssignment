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

	req, err := http.NewRequest("POST", "/persons/", bytes.NewBuffer(jsonReq))
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
	// expected := `{"id":1,"name":"javed","age":25,"email":"javed@tftus.com"}`
	// if r.Body.String() != expected {
	// 	t.Errorf("Wrong Answer By Handler: got %v want %v",
	// 		r.Body.String(), expected)
	// }
}

//  test wheather update work correctly or not...
func TestUpdatePersonByID(t *testing.T) {

	var jsonReq = []byte(`{"id":2,"name":"rajat","age":25,"email":"rajat@tftus.com"}`)

	req, err := http.NewRequest("PUT", "/persons/{id}", bytes.NewBuffer(jsonReq))
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
	// expected := `{"id":2,"name":"rajat","age":25,"email":"rajat@tftus.com"}`
	// if r.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		r.Body.String(), expected)
	// }
}

//  here testing getall func
func TestGetAll(t *testing.T) {

	req, err := http.NewRequest("GET", "/persons/", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(GetAllPerson)
	control.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}

	// expected := `[{id":1,"name":"javed","age":25,"email":"javed@tftus.com"},{"name":"rajat","age":25,"email":"rajat@tftus.com"}]`
	// if r.Body.String() != expected {
	// 	t.Errorf("Wrong Answer By Handler: got %v want %v",
	// 		r.Body.String(), expected)
	// }
}

//  testing getbyid func
func TestGetByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/persons/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(GetPersonById)
	control.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}

	// expected := `{"id":1,"name":"javed","age":25,"email":"javed@tftus.com"}`
	// if r.Body.String() != expected {
	// 	t.Errorf("Wrong Answer By Handler: got %v want %v",
	// 		r.Body.String(), expected)
	// }
}

func TestDeletPersonByID(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/persons/2", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(DeletePerson)
	control.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}

	// expected := `{"id":1,"name":"javed","age":25,"email":"javed@tftus.com"}`
	// if r.Body.String() != expected {
	// 	t.Errorf("Wrong Answer By Handler: got %v want %v",
	// 		r.Body.String(), expected)
	// }

}
