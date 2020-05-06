package main

import (
	"go-postgresql-userdb/internal/repositories"
	"go-postgresql-userdb/internal/web"
	"log"
)

func main() {
	err := repositories.CreateDB()
	if err != nil {
		log.Fatal("error on repositories.CreateDB: ", err)
	}
	web.Run()
}
