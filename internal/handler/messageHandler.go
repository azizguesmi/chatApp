package handler

import (
	"backend/internal/service"
	"encoding/json"
	"net/http"
)

type reqMessage struct {
	Content       string `json:"content"`
	ReceiverId    int    `json:"receiver_id"`
	Receiver_type string `json:"receiver_type"`
}

func HandelAddMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var req reqMessage
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := service.GetUserById(r.Context().Value("userID").(int))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	receiver, err := service.GetUserById(req.ReceiverId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if receiver == nil {
		http.Error(w, "Receiver not found", http.StatusNotFound)
		return
	}
	_, err = service.AddMessage(req.Content, user.ID, receiver.ID, req.Receiver_type)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "message added successfully"})
}

type delReq struct {
	IDMessage int `json:"id_message"`
}

func HandelDeleteMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	user, err := service.GetUserById(r.Context().Value("userID").(int))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	var req delReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	test, err := service.RemoveMessage(req.IDMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !test {
		http.Error(w, "message not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "message deleted"})
}

type updateReq struct {
	MessageId int `json:"id"`
	Content string `json:"content"`
}

func HandelUpdateMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost{
		http.Error(w, "methode not Allowed", http.StatusMethodNotAllowed)
		return
	}
	user, err := service.GetUserById(r.Context().Value("userID").(int))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	var req updateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	test, err := service.UpdateMessageContent(req.MessageId, req.Content, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !test {
		http.Error(w, "message not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "message updated"})
}







