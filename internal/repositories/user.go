package repositories

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go-postgresql-userdb/internal/model"
)

const (
	DB_USER = "admin"
	DB_NAME = "userinfo"
	DB_PASS = ""
)

var db *sqlx.DB

type User interface {
	AddUser(tx *sql.Tx, usr model.User) error
	DeleteUser(tx *sql.Tx, id int) error
	GetUserList(tx *sql.Tx) ([]model.User, error)
	UpdateUser(tx *sql.Tx, id int, usr model.User) error
}

// NewUser creates a new User repository
func NewUser() User {
	return &user{}
}

type user struct{}

func OpenDB() (*sqlx.DB, error) {
	dbinfo := fmt.Sprintf("postgres://%s:@localhost:5432/%s?sslmode=disable",
		DB_USER, DB_NAME)
	db, err := sqlx.Open("postgres", dbinfo)
	if err != nil {
		return nil, fmt.Errorf("error on connecting to db: %s", err)
	}
	return db, nil
}

func (u *user) AddUser(tx *sql.Tx, usr model.User) error {
	query := "INSERT INTO users (id, name, lastname, age, birthdate) VALUES($1, $2, $3, $4, $5)"

	_, err := tx.Exec(query, usr.Id, usr.Name, usr.Lastname, usr.Age, usr.Birthdate)
	if err != nil {
		return err
	}
	return nil
}

func (u *user) GetUserList(tx *sql.Tx) ([]model.User, error) {
	var users []model.User
	query := "SELECT * FROM users;"
	rows, err := tx.Query(query)
	if err != nil {
		return nil, err
	}

	users, err = processRows(rows)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *user) DeleteUser(tx *sql.Tx, id int) error {
	query := "DELETE FROM users WHERE id=$1"

	_, err := tx.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *user) UpdateUser(tx *sql.Tx, id int, usr model.User) error {
	query := "UPDATE users SET id=$1, name=$2, lastname=$3, age=$4, birthdate=$5 WHERE id=$6"

	_, err := tx.Exec(query, usr.Id, usr.Name, usr.Lastname, usr.Age, usr.Birthdate, id)
	if err != nil {
		return err
	}
	return nil
}

func processRows(rows *sql.Rows) ([]model.User, error) {
	defer rows.Close()

	users := make([]model.User, 0)

	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Lastname, &user.Age, &user.Birthdate)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	err := rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
