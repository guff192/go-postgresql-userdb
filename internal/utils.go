package internal

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const (
	DB_USER = "admin"
	DB_NAME = "userinfo"
	DB_PASS = ""
)

var db *sqlx.DB
var u struct {
	Id        int
	Name      string
	Lastname  string
	Age       int
	Birthdate string
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	checkError(err)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&u)
	checkError(err)

	dbinfo := fmt.Sprintf("postgres://%s:@localhost:5432/%s?sslmode=disable",
		DB_USER, DB_NAME)
	db, err := sqlx.Open("postgres", dbinfo)
	checkError(err)
	defer db.Close()
	_, err = db.Exec("INSERT INTO users VALUES($1, $2, $3, $4, $5)",
		u.Id, u.Name, u.Lastname, u.Age, u.Birthdate)
	checkError(err)
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	var Users []User
	dbinfo := fmt.Sprintf("postgres://%s:@localhost:5432/%s?sslmode=disable",
		DB_USER, DB_NAME)
	db, err := sqlx.Open("postgres", dbinfo)
	checkError(err)
	defer db.Close()
	err = db.Select(&Users, "SELECT * FROM users")
	checkError(err)

	w.Header().Set("Content-Type", "application/json") //setting content-type to JSON
	json_bytes, err := json.MarshalIndent(Users, "", "  ")
	checkError(err)
	fmt.Fprint(w, string(json_bytes))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	dbinfo := fmt.Sprintf("postgres://%s:@localhost:5432/%s?sslmode=disable",
		DB_USER, DB_NAME)
	db, err := sqlx.Open("postgres", dbinfo)
	checkError(err)
	defer db.Close()
	_, err = db.Exec("DELETE FROM users WHERE id=$1", id)
	checkError(err)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := r.ParseForm()
	checkError(err)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&u)
	checkError(err)

	dbinfo := fmt.Sprintf("postgres://%s:@localhost:5432/%s?sslmode=disable",
		DB_USER, DB_NAME)
	db, err := sqlx.Open("postgres", dbinfo)
	checkError(err)
	defer db.Close()

	_, err = db.Exec("UPDATE users SET id=$1, name=$2, lastname=$3, age=$4, birthdate=$5 WHERE id=$6",
		u.Id, u.Name, u.Lastname, u.Age, u.Birthdate, id)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
