package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-postgresql-userdb/internal/repositories"
	"go-postgresql-userdb/internal/web/binders"
	"net/http"
	"strconv"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		message := fmt.Sprint("error on parsing form: ", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
	user, err := binders.ParseUser(r.Body)
	if err != nil {
		message := fmt.Sprint("error on parsing user data from form: ", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	err = repositories.AddUser(user)
	if err != nil {
		message := fmt.Sprint("error on adding user to db: ", err)
		http.Error(w, message, http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Successfully added user to db."))
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	Users, err := repositories.GetUsers()
	if err != nil {
		message := fmt.Sprint("error on repositories.GetUsers: ", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") //setting content-type to JSON
	if err := binders.EncodeUsers(w, Users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		message := fmt.Sprint("missing or bad id: ", err)
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	if err = repositories.DeleteUser(id); err != nil {
		message := fmt.Sprint("error on repositories.DeleteUser: ", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully deleted user from db."))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		message := fmt.Sprint("missing or bad id: ", err)
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		message := fmt.Sprint("error on parsing form: ", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
	user, err := binders.ParseUser(r.Body)
	if err != nil {
		message := fmt.Sprint("error on parsing user data from form: ", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	if err = repositories.UpdateUser(id, user); err != nil {
		message := fmt.Sprint("error on repositories.DeleteUser: ", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully updated user."))
}
