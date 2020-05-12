package startup

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	ConfigDirectory = "./config"
	ConfigFile      = "startup.json"
)

type IniData struct {
	Port         int
	DBName       string
	DBHost       string
	DBPort       int
	UserName     string
	UserPassword string
}

func Configuration() (*IniData, error) {
	data := &IniData{}
	dirPath, err := filepath.Abs(ConfigDirectory)
	if err != nil {
		return data, err
	}
	path := filepath.Join(dirPath, ConfigFile)

	file, err := os.Open(path)
	if err != nil {
		return data, err
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&data); err != nil {
		return data, err
	}
	return data, nil
}
