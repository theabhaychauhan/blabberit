package user

import (
	"encoding/json"
	"net/http"
)

// ------- Request DTO --------
type registerRequest struct {
	Username  string `json:"username"`
	PublicKey string `json:"publickey"`
}

// --------- Handlers ---------
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req registerRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Username == "" || req.PublicKey == "" {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if existing, _ := FindUserByUsername(req.Username); existing != nil {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	if err := CreateUser(&User{
		Username:  req.Username,
		PublicKey: req.PublicKey,
	}); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User Registered",
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is Required", http.StatusBadRequest)
		return
	}

	user, err := FindUserByUsername(username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{
		"username":  user.Username,
		"publicKey": user.PublicKey,
	})
}
