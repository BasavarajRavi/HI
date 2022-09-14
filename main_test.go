package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_display(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/get", nil)
	display(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status: %d, but got : %d", http.StatusOK, w.Code)
	}

}
