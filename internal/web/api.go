package web

import (
	"github.com/gorilla/mux"
	"go-postgresql-userdb/internal/web/controllers"
	"log"
	"net/http"
)

func Run() {
	router := mux.NewRouter()
	handleRoutes(router)

	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("unexpected error while ListenAndServe: ", err)
	}
}

func handleRoutes(router *mux.Router) {
	router.HandleFunc("/users", controllers.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/users", controllers.GetUserList).Methods(http.MethodGet)
	router.HandleFunc("/users/:{id:[0-9]+}", controllers.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/users/:{id:[0-9]+}", controllers.UpdateUser).Methods(http.MethodPut)
}
