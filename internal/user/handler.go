package user

import (
	"encoding/json"
	"net/http"
)

type registerRequest struct {
	Username  string `json:"username"`
	PublicKey string `json:"publickey"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req registerRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Username == "" || req.PublicKey == "" {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	var existing User
	if err := DB.Where("username = ?", req.Username).First(&existing).Error; err == nil {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	user := User{
		Username:  req.Username,
		PublicKey: req.PublicKey,
	}

	result := DB.Create(&user)
	if result.Error != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User Registered",
	})
}
