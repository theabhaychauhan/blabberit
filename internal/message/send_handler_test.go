package message

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/thechauhanabhay/blabberit/internal/testutils"
	"github.com/thechauhanabhay/blabberit/internal/user"
)

func setupMsgTestDB(t *testing.T) {
	t.Helper()

	_ = godotenv.Load("../../.env.test")
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		t.Fatal("POSTGRES_DSN not set")
	}

	db := testutils.SetupTestDB(dsn, &user.User{}, &Message{})
	user.SetDB(db)
	SetDB(db)
}

func TestSendMessageHandler_Success(t *testing.T) {
	setupMsgTestDB(t)

	body := `{"from":"keyA","to":"keyB","content":"hello UTC"}`
	req := httptest.NewRequest(http.MethodPost, "/send", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	SendMessageHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}

	var resp map[string]any
	_ = json.NewDecoder(w.Body).Decode(&resp)
	if resp["message"] != "Message Sent" {
		t.Errorf("unexpected response body: %v", resp)
	}

	msgs, err := GetMessagesForUser("keyB")
	if err != nil || len(msgs) != 1 {
		t.Fatalf("expected 1 saved message, got %d (err=%v)", len(msgs), err)
	}

	if !msgs[0].Timestamp.Equal(msgs[0].Timestamp.UTC()) {
		t.Errorf("timestamp not saved in UTC: %v", msgs[0].Timestamp)
	}
}

func TestSendMessageHandler_InvalidRequest(t *testing.T) {
	setupMsgTestDB(t)

	req := httptest.NewRequest(http.MethodPost, "/send",
		bytes.NewBufferString(`{"from":"","to":"x","content":""}`))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	SendMessageHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}
