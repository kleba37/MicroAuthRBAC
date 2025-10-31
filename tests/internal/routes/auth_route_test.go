package routes_test

import (
	"bytes"
	"encoding/json"
	MessageResponses "go-test/api/Responses"
	TestingTools "go-test/pkg/Testing"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "modernc.org/sqlite"
)

func setupTest(t *testing.T) *TestingTools.TestingTools {
	testingTools, err := TestingTools.New()
	if err != nil {
		t.Fatal(err)
	}
	return testingTools
}

func TestAuthRoute(t *testing.T) {
	tt := setupTest(t)

	_, err := tt.DB.Exec(`INSERT INTO users (name, email, token) VALUES (?, ?, ?)`, "Test User", "test@example.com", "secret")
	if err != nil {
		t.Fatal(err)
	}

	err = tt.StartTest()
	if err != nil {
		t.Fatal(err)
	}

	// Сценарий 1: Валидный токен
	t.Run("Valid operation token should return 200 OK", func(t *testing.T) {
		body, _ := json.Marshal(map[string]string{"token": "secret"})
		req, _ := http.NewRequest("POST", "/access", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer token")

		rr := httptest.NewRecorder()
		tt.Serve(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var response MessageResponses.AuthOperationResponse
		if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
			t.Fatalf("Failed to unmarshal response body: %v", err)
		}

		if !response.Access || response.Message != "Success" {
			t.Errorf("handler returned unexpected body: got %+v", response)
		}
	})

	t.Run("Invalid operation token should return 403 Forbidden", func(t *testing.T) {
		body, _ := json.Marshal(map[string]string{"token": "wrong_secret"})
		req, _ := http.NewRequest("POST", "/access", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer token")

		rr := httptest.NewRecorder()
		tt.Serve(rr, req)

		if status := rr.Code; status != http.StatusForbidden {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusForbidden)
		}

		var response MessageResponses.AuthOperationResponse
		if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
			t.Fatalf("Failed to unmarshal response body: %v", err)
		}

		if response.Status != http.StatusForbidden || response.Message != "Not Authorized" {
			t.Errorf("handler returned unexpected body: got %+v", response)
		}
	})
}
