package services

import (
	"go-postgresql-userdb/internal/model"
	"go-postgresql-userdb/internal/repositories"
)

type User interface {
	AddUser(user model.User) error
	DeleteUser(id int) error
	GetUserList() ([]model.User, error)
	UpdateUser(id int, user model.User) error
}

// NewUser creates User service
func NewUser(repository repositories.User) User {
	return &user{repository: repository}
}

type user struct {
	repository repositories.User
}

func (u *user) AddUser(user model.User) error {
	if err := u.repository.AddUser(user); err != nil {
		return err
	}
	return nil
}

func (u *user) GetUserList() ([]model.User, error) {
	users, err := u.repository.GetUserList()
	if err != nil {
		return []model.User{}, err
	}
	return users, nil
}

func (u *user) DeleteUser(id int) error {
	if err := u.repository.DeleteUser(id); err != nil {
		return err
	}
	return nil
}

func (u *user) UpdateUser(id int, user model.User) error {
	if err := u.repository.UpdateUser(id, user); err != nil {
		return err
	}
	return nil
}
