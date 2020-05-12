package services

import (
	"go-postgresql-userdb/internal/datasource"
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
	tx, err := datasource.SQL.Begin()
	if err != nil {
		return err
	}
	defer func() {
		err = datasource.CloseTransaction(tx, err)
	}()

	if err := u.repository.AddUser(tx, user); err != nil {
		return err
	}
	return nil
}

func (u *user) GetUserList() ([]model.User, error) {
	tx, err := datasource.SQL.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = datasource.CloseTransaction(tx, err)
	}()

	users, err := u.repository.GetUserList(tx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *user) DeleteUser(id int) error {
	tx, err := datasource.SQL.Begin()
	if err != nil {
		return err
	}
	defer func() {
		err = datasource.CloseTransaction(tx, err)
	}()

	if err := u.repository.DeleteUser(tx, id); err != nil {
		return err
	}
	return nil
}

func (u *user) UpdateUser(id int, user model.User) error {
	tx, err := datasource.SQL.Begin()
	if err != nil {
		return err
	}
	defer func() {
		err = datasource.CloseTransaction(tx, err)
	}()

	if err := u.repository.UpdateUser(tx, id, user); err != nil {
		return err
	}
	return nil
}
