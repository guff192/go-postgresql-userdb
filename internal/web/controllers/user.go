package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"go-postgresql-userdb/internal/model"
	"go-postgresql-userdb/internal/services"
	"go-postgresql-userdb/internal/utils"
	"go-postgresql-userdb/internal/web/binders"
	"io"
	"net/http"
)

type User interface {
	AddUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
	GetUserList(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
}

// NewUser produces a user's controller
func NewUser(service services.User) User {
	return &user{service: service}
}

type user struct {
	service services.User
}

func (u *user) AddUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		message := fmt.Sprint("error on parsing form: ", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
	user, err := parseUser(r.Body)
	if err != nil {
		message := fmt.Sprint("error on parsing user data from form: ", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	if err = u.service.AddUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Successfully added user to db."))
}

func (u *user) GetUserList(w http.ResponseWriter, r *http.Request) {
	Users, err := u.service.GetUserList()
	if err != nil {
		message := fmt.Sprint("error on repositories.GetUserList: ", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") //setting content-type to JSON
	if err := encodeUsers(w, Users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (u *user) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := context.Get(r, binders.ID).(int)

	if err := u.service.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully deleted user."))
}

func (u *user) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := context.Get(r, binders.ID).(int)

	user, err := parseUser(r.Body)
	if err != nil {
		message := fmt.Sprint("error on parsing user data from form: ", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	if err = u.service.UpdateUser(id, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully updated user."))
}

func parseUser(rawData io.Reader) (model.User, error) {
	var user struct {
		Id        int
		Name      string
		Lastname  string
		Age       int
		Birthdate string
	}
	decoder := json.NewDecoder(rawData)
	err := decoder.Decode(&user)
	if err != nil {
		return model.User{}, fmt.Errorf("error on jsonDecode: %s", err)
	}
	birthdate, err := utils.ParseDate(user.Birthdate)
	if err != nil {
		return model.User{}, fmt.Errorf("error on parsing date: %s", err)
	}
	usr := model.User{
		Id:        user.Id,
		Name:      user.Name,
		Lastname:  user.Lastname,
		Age:       user.Age,
		Birthdate: *birthdate,
	}
	return usr, nil
}

func encodeUsers(w io.Writer, users interface{}) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(users); err != nil {
		return fmt.Errorf("error on jsonEncode: %s", err)
	}
	return nil
}
