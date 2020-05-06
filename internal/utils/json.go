package utils

import (
	"encoding/json"
	"go-postgresql-userdb/internal/model"
	"io"
	"strconv"
	"strings"
	"time"
)

var user struct {
	Id        int
	Name      string
	Lastname  string
	Age       int
	Birthdate string
}

func ParseUser(rawData io.Reader) (*model.User, error) {
	decoder := json.NewDecoder(rawData)
	err := decoder.Decode(&user)
	if err != nil {
		return nil, err
	}
	birthdate, err := parseDate(user.Birthdate)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Id:        user.Id,
		Name:      user.Name,
		Lastname:  user.Lastname,
		Age:       user.Age,
		Birthdate: *birthdate,
	}
	return user, nil
}

func EncodeUsers(users []model.User) (string, error) {
	jsonBytes, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return "", err
	}
	jsonString := string(jsonBytes)
	return jsonString, nil
}

func parseDate(sDate string) (*time.Time, error) {
	dateSl := strings.Split(sDate, "-")
	y, err := strconv.Atoi(dateSl[0])
	if err != nil {
		return nil, err
	}

	m, err := strconv.Atoi(dateSl[1])
	if err != nil {
		return nil, err
	}

	d, err := strconv.Atoi(dateSl[2])
	if err != nil {
		return nil, err
	}

	date := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
	return &date, nil
}
