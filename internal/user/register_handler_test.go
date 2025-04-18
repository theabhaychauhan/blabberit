package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/thechauhanabhay/blabberit/internal/testutils"
)

func setupTestDB() {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		panic("Could not load .env.test")
	}

	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		panic("POSTGRES_DSN not set")
	}

	testutils.DB = testutils.SetupTestDB(dsn, &User{})
	DB = testutils.DB
}

func TestRegisterHandler_Success(t *testing.T) {
	setupTestDB()

	payload := registerRequest{
		Username:  "alice",
		PublicKey: "pubkey-alice",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	RegisterHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201 Created, got %d", w.Code)
	}
}

func TestRegisterHandler_InvalidRequest(t *testing.T) {
	setupTestDB()

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	RegisterHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 Bad Request, got %d", w.Code)
	}
}

func TestRegisterHandler_DuplicateUsername(t *testing.T) {
	setupTestDB()

	payload := registerRequest{
		Username:  "bob",
		PublicKey: "pubkey-bob",
	}
	body, _ := json.Marshal(payload)

	req1 := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	RegisterHandler(w1, req1)

	req2 := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	RegisterHandler(w2, req2)

	if w2.Code != http.StatusConflict {
		t.Errorf("expected 409 Conflict, got %d", w2.Code)
	}
}
