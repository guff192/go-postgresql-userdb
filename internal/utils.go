package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	decoder := json.NewDecoder(r.Body)
	var u User
	err := decoder.Decode(&u)
	checkError(err)
	Users = append(Users, u)
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //setting content-type to JSON
	json_bytes, err := json.MarshalIndent(Users, "", "  ")
	checkError(err)
	fmt.Fprint(w, string(json_bytes))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
