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

type reqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandelAddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
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

func HandelLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	var req reqLogin
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := service.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{
		"user_id": int64(user.ID),
	})
}

func HandelDeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	type req struct {
		Id int `json:"id"`
	}
	var re req
	if err := json.NewDecoder(r.Body).Decode(&re); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	test, err := service.DeleteUser(re.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var res string
	if !test {
		res = "user not found"
	} else {
		res = "Deleted succesfully"
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": res,
	})
}
