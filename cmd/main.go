package main

import (
	"go-postgresql-userdb/internal/datsource"
	"go-postgresql-userdb/internal/web"
	"log"
)

func main() {
	err := datsource.CreateDB()
	if err != nil {
		log.Fatal("error on repositories.CreateDB: ", err)
	}
	web.Run()
}
