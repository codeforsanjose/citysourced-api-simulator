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

func Init(fileName string) error {
	log.Info("Loading configuration file - Config: %q", fileName)
	_, err := readConfig(fileName)
	if err != nil {
		return errors.New(fmt.Sprintf("Error loading config: %s", err))
	}
	return nil
}

func Auth(ac string) bool {
	if ac == C.API.AuthKey {
		return true
	}
	return false
}

// ==============================================================================================================================
//                                      CONFIG
// ==============================================================================================================================

// ------------------------------- ConfigType -------------------------------
type ConfigType struct {
	Loaded          bool
	Instrumentation DebugType `json:"instrumentation"`
	API             API_Type  `json:"api"`
}

func (x *ConfigType) Display() string {
	s := fmt.Sprintf("\n==================================== CONFIG ==================================\n")
	s += spew.Sdump(x)
	s += fmt.Sprintf("==============================================================================\n")
	return s
}

// ------------------------------- DebugType -------------------------------
type DebugType struct {
	Debug   bool `json:"debug"`
	Verbose bool `json:"verbose"`
}

// ------------------------------- API_Type -------------------------------
type API_Type struct {
	AuthKey string `json:"authkey"`
}

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
