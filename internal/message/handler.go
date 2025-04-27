package message

import (
	"encoding/json"
	"net/http"
	"time"
)

type SendMessageRequest struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var req SendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil ||
		req.From == "" || req.To == "" || req.Content == "" {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	message := Message{
		From:      req.From,
		To:        req.To,
		Content:   req.Content,
		Timestamp: time.Now().UTC(),
	}

	if err := SaveMessage(&message); err != nil {
		http.Error(w, "Failed to save message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"id":       message.ID,
		"message":  "Message Sent",
		"datetime": message.Timestamp,
	})
}

func FetchMessagesHandler(w http.ResponseWriter, r *http.Request) {
	pubKey := r.URL.Query().Get("user")
	if pubKey == "" {
		http.Error(w, "user query-param required", http.StatusBadRequest)
		return
	}

	messages, err := GetMessagesForUser(pubKey)
	if err != nil {
		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(messages)
}
