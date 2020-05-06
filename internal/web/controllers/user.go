package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-postgresql-userdb/internal/repositories"
	"go-postgresql-userdb/internal/utils"
	"log"
	"net/http"
	"strconv"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	checkError(err)
	user, err := utils.ParseUser(r.Body)
	checkError(err)

	err = repositories.AddUser(user)
	checkError(err)
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	Users, err := repositories.GetUsers()
	checkError(err)

	w.Header().Set("Content-Type", "application/json") //setting content-type to JSON
	jsonUsers, err := utils.EncodeUsers(*Users)
	checkError(err)
	fmt.Fprint(w, jsonUsers)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	checkError(err)

	err = repositories.DelUser(id)
	checkError(err)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	checkError(err)

	err = r.ParseForm()
	checkError(err)
	user, err := utils.ParseUser(r.Body)
	checkError(err)

	err = repositories.UpdUser(id, user)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
