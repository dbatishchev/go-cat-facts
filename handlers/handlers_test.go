package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MainPageHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), `<h1 class="cover-heading">`) {
		t.Errorf("handler returned unexpected body: got %v want html output with header",
			rr.Body.String())
	}
}
