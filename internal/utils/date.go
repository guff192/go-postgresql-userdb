package utils

import (
	"strconv"
	"strings"
	"time"
)

func ParseDate(sDate string) (*time.Time, error) {
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
