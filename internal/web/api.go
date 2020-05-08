package web

import (
	"github.com/gorilla/mux"
	"go-postgresql-userdb/internal/repositories"
	"go-postgresql-userdb/internal/services"
	"go-postgresql-userdb/internal/web/binders"
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
	controller := newController()

	router.HandleFunc("/users", controller.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/users", controller.GetUserList).Methods(http.MethodGet)
	router.HandleFunc("/users/:{id:[0-9]+}",
		binders.IDBinder(controller.DeleteUser)).Methods(http.MethodDelete)
	router.HandleFunc("/users/:{id:[0-9]+}",
		binders.IDBinder(controller.UpdateUser)).Methods(http.MethodPut)
}

func newController() controllers.User {
	repository := repositories.NewUser()
	service := services.NewUser(repository)
	return controllers.NewUser(service)
}
