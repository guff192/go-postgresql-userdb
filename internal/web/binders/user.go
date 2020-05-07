package binders

import (
	"encoding/json"
	"fmt"
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

func ParseUser(rawData io.Reader) (model.User, error) {
	decoder := json.NewDecoder(rawData)
	err := decoder.Decode(&user)
	if err != nil {
		return model.User{}, fmt.Errorf("error on jsonDecode: %s", err)
	}
	birthdate, err := parseDate(user.Birthdate)
	if err != nil {
		return model.User{}, fmt.Errorf("error on parsing date: %s", err)
	}
	user := model.User{
		Id:        user.Id,
		Name:      user.Name,
		Lastname:  user.Lastname,
		Age:       user.Age,
		Birthdate: *birthdate,
	}
	return user, nil
}

func EncodeUsers(w io.Writer, users interface{}) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(users); err != nil {
		return fmt.Errorf("error on jsonEncode: %s", err)
	}
	return nil
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
