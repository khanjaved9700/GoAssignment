package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	DataMigration()
}

func TestCreatePerson(t *testing.T) {

	var jsonReq = []byte(`{"id":1,"name":"javed","age":25,"email":"javed@tftus.com"}`)

	req, err := http.NewRequest("POST", "/persons", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	r := httptest.NewRecorder()

	control := http.HandlerFunc(createPerson)
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

func TestUpdatePersonByID(t *testing.T) {

	var jsonReq = []byte(`{"id":2,"name":"rajat","age":25,"email":"rajat@tftus.com"}`)

	req, err := http.NewRequest("PUT", "/persons/{id}", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	r := httptest.NewRecorder()
	control := http.HandlerFunc(updateById)
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

func TestGetAll(t *testing.T) {

	req, err := http.NewRequest("GET", "/persons", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(getAll)
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

func TestGetByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/persons/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(getById)
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
	control := http.HandlerFunc(deleteById)
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
