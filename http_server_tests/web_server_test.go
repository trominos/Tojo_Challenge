package main

import (
	// "fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPtoHTTPSRedirect(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://localhost/", http.StatusMovedPermanently)
	})
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMovedPermanently {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusMovedPermanently)
	}
}

