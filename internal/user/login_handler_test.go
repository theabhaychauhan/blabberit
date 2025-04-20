package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/thechauhanabhay/blabberit/internal/testutils"
)

func setupLoginTestDB() {
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

func TestLoginHandler_UserFound(t *testing.T) {
	setupLoginTestDB()

	DB.Create(&User{
		Username:  "abhay",
		PublicKey: "Testing-Testing",
	})

	req := httptest.NewRequest("GET", "/login?username=abhay", nil)
	w := httptest.NewRecorder()

	LoginHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200, got %d", w.Code)
	}

	var resp map[string]string
	println("BODY:", w.Body.String())

	json.NewDecoder(w.Body).Decode(&resp)

	if resp["publicKey"] != "Testing-Testing" {
		t.Errorf("Expected publickey to be 'Testing-Testing', got %s", resp["publickey"])
	}
}

func TestLoginHandler_UserNotFound(t *testing.T) {
	setupLoginTestDB()

	req := httptest.NewRequest("GET", "/login?username=ghost", nil)
	w := httptest.NewRecorder()

	LoginHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", w.Code)
	}
}

func TestLoginHandler_MissingUsername(t *testing.T) {
	setupLoginTestDB()

	req := httptest.NewRequest("GET", "/login", nil)
	w := httptest.NewRecorder()

	LoginHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}
