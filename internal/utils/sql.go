package utils

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func GetScript(dirname, filename string) ([]string, error) {
	dirname, err := filepath.Abs(dirname)
	if err != nil {
		return nil, err
	}
	path := filepath.Join(dirname, filename)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	sData := strings.TrimRight(string(data), ";\n")
	return strings.Split(sData, ";"), nil
}
