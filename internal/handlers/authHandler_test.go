package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func testLogin(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AuthHandler)
}
