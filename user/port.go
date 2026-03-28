package user

import (
	"sysagent/domain"
	userHandler "sysagent/rest/handlers/user"
)

type Service interface {
	userHandler.UserService //embedding
}

type UserRepo interface {
	Create(p domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
	// List() ([]*User, error)
	// Delete(userID int) error
	// Update(user User) (*User, error)
}
