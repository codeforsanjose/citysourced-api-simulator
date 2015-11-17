package config

import (
	"CitySourcedAPI/logs"

	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
)

var (
	log = logs.Log
	C   ConfigType
)

func Init(fileName string, port int64) error {
	log.Info("Loading config: %q", fileName)
	_, err := readConfig(fileName)
	if err != nil {
		return errors.New(fmt.Sprintf("Error loading config: %s", err))
	}
	if port > 0 {
		C.Server.Port = port
	}
	return nil
}

func Auth(ac string) bool {
	if ac == C.API.AuthKey {
		log.Debug("Auth OK - req: %q  key: %q", ac, C.API.AuthKey)
		return true
	}
	log.Debug("Auth FAIL - req: %q  key: %q", ac, C.API.AuthKey)
	return false
}

func Port() int64 {
	return C.Server.Port
}

// ==============================================================================================================================
//                                      CONFIG
// ==============================================================================================================================

func readConfig(filePath string) (*ConfigType, error) {
	if C.Loaded {
		msg := "Duplicate calls to load Config file!"
		log.Warning(msg)
		return &C, errors.New(msg)
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		msg := fmt.Sprintf("Unable to open the Config file: %q.\nError: %v", filePath, err)
		log.Critical(msg)
		return nil, errors.New(msg)
	}

	err = json.Unmarshal([]byte(file), &C)
	if err != nil {
		msg := fmt.Sprintf("Invalid JSON in the Config file: %q.\nError: %v", filePath, err)
		log.Critical(msg)
		return nil, errors.New(msg)
	}

	C.Loaded = true
	return &C, nil
}

// ------------------------------- ConfigType -------------------------------
type ConfigType struct {
	Loaded bool
	Server ServerType `json:"server"`
	API    API_Type   `json:"api"`
}

func (x *ConfigType) Display() string {
	s := fmt.Sprintf("\n==================================== CONFIG ==================================\n")
	s += spew.Sdump(x)
	s += fmt.Sprintf("==============================================================================\n")
	return s
}

// ------------------------------- API_Type -------------------------------
type API_Type struct {
	AuthKey string `json:"authkey"`
}

// ------------------------------- ServerType -------------------------------
type ServerType struct {
	Port int64 `json:"port"`
}
