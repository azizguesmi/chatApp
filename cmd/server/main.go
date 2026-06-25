package main

import (
	"backend/internal/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/register", handler.HandelAddUser)
	mux.HandleFunc("/auth/login", handler.HandelLogin)
	mux.HandleFunc("/auth/delete-user", handler.HandelDeleteUser)

	http.ListenAndServe(":8080", mux)
}
