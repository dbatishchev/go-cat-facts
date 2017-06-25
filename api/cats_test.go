package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
)

func TestMakeRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"fake twitter json string"}`)
	}))
	defer ts.Close()

	result, err := makeRequest(ts.URL)

	if err != nil {
		t.Errorf("Failed fetch response from API: %v", err)
	}

	if len(result) == 0 {
		t.Error("Empty response from API")
	}
}

func TestGetFactFromJSON(t *testing.T) {
	fact, err := getFactFromJSON([]byte(`{"facts": ["fact1", "fact2"]}`))

	if err != nil {
		t.Errorf("Failed to decode response JSON: %v", err)
	}

	if len(fact) == 0 {
		t.Error("Failed to decode response JSON")
	}
}