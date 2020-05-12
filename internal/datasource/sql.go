package datasource

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-postgresql-userdb/internal/init/startup"
	"go-postgresql-userdb/internal/utils"
)

const (
	ScriptsDirectory = "./scripts"
	CreateScript     = "schema.sql"
)

var SQL *sqlx.DB

func InitSQL(iniData *startup.IniData) error {
	if err := connect(iniData); err != nil {
		return err
	}

	if err := createTable(); err != nil {
		return err
	}

	return nil
}

func connect(iniData *startup.IniData) error {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		iniData.UserName, iniData.UserPassword, iniData.DBHost, iniData.DBPort, iniData.DBName)

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return fmt.Errorf("error on connecting to db: %s", err)
	}

	SQL = db
	return nil
}

func createTable() (err error) {
	tx, err := SQL.Begin()
	if err != nil {
		return err
	}
	defer func() {
		err = CloseTransaction(tx, err)
	}()

	queries, err := utils.GetScript(ScriptsDirectory, CreateScript)
	if err != nil {
		return err
	}

	for _, query := range queries {
		_, err = tx.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

func CloseTransaction(tx *sql.Tx, err error) error {
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			err = fmt.Errorf("%w: %q", err, rbErr)
		}
		return err
	}

	cmtErr := tx.Commit()
	if cmtErr != nil {
		err = cmtErr
	}
	return err
}
