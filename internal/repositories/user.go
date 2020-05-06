package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go-postgresql-userdb/internal/model"
	"go-postgresql-userdb/internal/utils"
)

const (
	DB_USER          = "admin"
	DB_NAME          = "userinfo"
	DB_PASS          = ""
	ScriptsDirectory = "./scripts"
	CreateScript     = "schema.sql"
)

type user struct{}

func CreateDB() error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	queries, err := utils.Script(ScriptsDirectory, CreateScript)
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

func openDB() (*sqlx.DB, error) {
	dbinfo := fmt.Sprintf("postgres://%s:@localhost:5432/%s?sslmode=disable",
		DB_USER, DB_NAME)
	db, err := sqlx.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func AddUser(u *model.User) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO users (id, name, lastname, age, birthdate) VALUES($1, $2, $3, $4, $5)"
	_, err = db.Exec(query, u.Id, u.Name, u.Lastname, u.Age, u.Birthdate)
	if err != nil {
		return err
	}
	return nil
}

func GetUsers() (*[]model.User, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var users []model.User
	query := "SELECT * FROM users;"
	err = db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func DelUser(id int) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdUser(id int, u *model.User) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "UPDATE users SET id=$1, name=$2, lastname=$3, age=$4, birthdate=$5 WHERE id=$6"
	_, err = db.Exec(query, u.Id, u.Name, u.Lastname, u.Age, u.Birthdate, id)
	if err != nil {
		return err
	}
	return nil
}
