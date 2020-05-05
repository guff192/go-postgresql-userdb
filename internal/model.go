package internal

import "time"

type User struct {
	Id        int
	Name      string
	Lastname  string
	Age       int
	Birthdate time.Time
}
