package handlers

import (
	"net/http"

	"github.com/raynine/ticket-system-backend/interfaces"
	"github.com/raynine/ticket-system-backend/repos"
)

type UserHandler struct {
	UserRepo interfaces.UserRepo
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		UserRepo: repos.NewUserRepo(),
	}
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
}
