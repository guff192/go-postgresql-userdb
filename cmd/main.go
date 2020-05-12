package main

import (
	"go-postgresql-userdb/internal/datasource"
	"go-postgresql-userdb/internal/init/startup"
	"go-postgresql-userdb/internal/web"
	"log"
)

func main() {
	iniData, err := startup.Configuration()
	if err != nil {
		log.Fatal("error on parsing config: ", err)
	}
	err = datasource.InitSQL(iniData)
	if err != nil {
		log.Fatal("error on datasource.InitSQL: ", err)
	}
	web.Run(iniData)
}
