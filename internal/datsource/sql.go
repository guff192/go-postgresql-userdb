package datsource

import (
	"go-postgresql-userdb/internal/repositories"
	"go-postgresql-userdb/internal/utils"
)

const (
	ScriptsDirectory = "./scripts"
	CreateScript     = "schema.sql"
)

func CreateDB() error {
	db, err := repositories.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	queries, err := utils.GetScript(ScriptsDirectory, CreateScript)
	if err != nil {
		return err
	}

	for _, query := range queries {
		_, err = db.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}
