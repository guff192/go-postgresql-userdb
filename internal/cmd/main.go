package main

import (
	"github.com/gorilla/mux"
	"go-postgresql-userdb/internal"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", internal.AddUser).Methods(http.MethodPost)
	r.HandleFunc("/users", internal.GetUserList).Methods(http.MethodGet)
	r.HandleFunc("/users/:{id:[0-9]+}", internal.DeleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/users/:{id:[0-9]+}", internal.UpdateUser).Methods(http.MethodPut)

	err := http.ListenAndServe(":9090", r)
	if err != nil {
		log.Fatal("Listen and serve: ", err)
	}

}
