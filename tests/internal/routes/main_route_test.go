package routes_test

import (
	"go-test/pkg/Router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainRoute(t *testing.T) {
	handler := http.HandlerFunc(Router.Router{}.Router)

	t.Run("Valid request to main route", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
		expected := `{"status":"ok","message":"Hello World!"}` + "\n"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})

	t.Run("Invalid request with non-existent path", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/invalid-path", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
}
