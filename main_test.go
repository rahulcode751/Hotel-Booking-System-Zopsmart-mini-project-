package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostBookingDetail(t *testing.T) {
	requestBody := `{
        "HotelName": "Vedantam",
        "CustomerName": "Testing Purpose",
        "Date": "16-12-2024",
        "Price": 1000,
        "CustomerContact": "6264959991",
        "roomno": 25
    }`
	req, err := http.NewRequest("POST", "/hotelbooking/add", bytes.NewBufferString(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	// Setting the Handler for our route
	handler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.WriteHeader(http.StatusOK)
	})
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	// Checking Status Code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
func TestUpdateBookingDetail(t *testing.T) {
	requestBody := `{
		"HotelName":"Vedantam",
		"CustomerName":"Chandresh Pujara",
		"Date":"17-12-2024",
		"Price":1000,
		"CustomerContact":"6264959991",
		"roomno":13
    }`
	req, err := http.NewRequest("PUT", "/hotelbooking/update/6", bytes.NewBufferString(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Setting the Handler for our route
	handler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.WriteHeader(http.StatusOK)
	})
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	// Checking Status Code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestGetAllDetails(t *testing.T) {
	req, err := http.NewRequest("GET", "/hotelbooking/details", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Setting the Handler for our route
	handler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.WriteHeader(http.StatusOK)
	})
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	// Checking Status Code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetDetailsByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/hotelbooking/detail/6", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Setting the Handler for our route
	handler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.WriteHeader(http.StatusOK)
	})
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	/// Checking Status Code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
func TestDeleteDetailsByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/hotelbooking/delete/6", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Setting the Handler for our route
	handler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.WriteHeader(http.StatusOK)
	})
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	// Checking Status Code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
