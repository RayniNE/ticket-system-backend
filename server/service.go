package handlers

import (
	"github.com/gorilla/mux"
	"github.com/raynine/ticket-system-backend/server/handlers"
)

func CreateServer() {
	router := mux.NewRouter()
	userHandler := handlers.NewUserHandler()
	InitializeUserHandler(router, userHandler)
}

func InitializeUserHandler(route *mux.Router, userHandler *handlers.UserHandler) {
	userRouter := route.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", userHandler.CreateUser).Methods("POST")
}
