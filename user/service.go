package user

import (
	"sysagent/domain"
)

type service struct {
	usrRepo UserRepo
}

func NewService(usrRepo UserRepo) Service {
	return &service{
		usrRepo: usrRepo,
	}
}

func (srv *service) Create(user domain.User) (*domain.User, error) {
	usr, err := srv.usrRepo.Create(user)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	return usr, nil
}

func (srv *service) Find(email string, pass string) (*domain.User, error) {
	usr, err := srv.usrRepo.Find(email, pass)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	return usr, nil
}
