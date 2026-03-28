package user

import (
	"sysagent/config"
)

type Handler struct {
	cnf *config.Config
	srv UserService
}

func NewHandler(cnf *config.Config, srv UserService) *Handler {
	return &Handler{
		cnf: cnf,
		srv: srv,
	}
}
