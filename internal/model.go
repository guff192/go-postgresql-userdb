package internal

import "time"

type Date struct {
	Day   int
	Month time.Month
	Year  int
}

type User struct {
	Id        int
	Name      string
	Lastname  string
	Age       int
	Birthdate Date
}

var Users []User
