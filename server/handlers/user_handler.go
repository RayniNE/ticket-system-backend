package handlers

import (
	"net/http"

	"github.com/raynine/ticket-system-backend/interfaces"
	"github.com/raynine/ticket-system-backend/models"
	"github.com/raynine/ticket-system-backend/repos"
)

type UserHandler struct {
	UserRepo interfaces.UserRepo
	Config   models.Config
}

func NewUserHandler(opts models.Config) *UserHandler {
	return &UserHandler{
		UserRepo: repos.NewUserRepo(),
		Config:   opts,
	}
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
}
