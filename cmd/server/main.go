package main

import (
	"fmt"
	"net/http"

	middleware "backend/internal/MiddleWare"
	"backend/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/register", handler.HandelAddUser)
	mux.HandleFunc("/auth/login", handler.HandelLogin)
	mux.HandleFunc("/auth/delete-user", handler.HandelDeleteUser)
	mux.HandleFunc("/message/send", middleware.AuthMiddleware(handler.HandelAddMessage))

	fmt.Println("server runing on port 8080")
	http.ListenAndServe(":8080", mux)
}
