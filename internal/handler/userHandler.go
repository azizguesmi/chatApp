package handler

import (
	"backend/internal/service"
	"encoding/json"
	"net/http"
)

type reqUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func HandelAddUser(w http.ResponseWriter, r *http.Request) {
	var req reqUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := service.AddUser(req.Username, req.Password, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{
		"message_id": id,
	})
}
