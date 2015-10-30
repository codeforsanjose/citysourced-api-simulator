package data

import (
	"CitySourcedAPI/common"
	"CitySourcedAPI/logs"

	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/davecgh/go-spew/spew"
)

var debug = true
var verbose = true
var log = logs.Log

// ==============================================================================================================================
//                                      CONFIG VARIABLES
// ==============================================================================================================================

var (
	System SystemType
)

func ReadConfig(filePathSystem string) error {
	log.Info("Loading configuration file - System: %q", filePathSystem)
	_, errSystem := readSystem(filePathSystem)
	if errSystem != nil {
		return errors.New("Configuration has already been loaded!")
	}
	return nil
}

func readSystem(filePath string) (*SystemType, error) {
	if System.Loaded {
		msg := "Duplicate calls to System Config!"
		log.Warning(msg)
		return &System, errors.New(msg)
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		msg := fmt.Sprintf("Unable to open the System Config file - specified at: %q.\nError: %v", filePath, err)
		log.Critical(msg)
		return nil, errors.New(msg)
	}

	err = json.Unmarshal([]byte(file), &System)
	if err != nil {
		msg := fmt.Sprintf("Invalid JSON in the System Config file - specified at: %q.\nError: %v", filePath, err)
		log.Critical(msg)
		return nil, errors.New(msg)
	}

	// Update Last ID
	var lastId int64 = 0
	for _, v := range System.Data.Reports {
		if v.Id > lastId {
			lastId = v.Id
		}
	}
	System.Data.lastId = lastId

	fmt.Printf("After loading dates...\n%s\n", spew.Sdump(System))

	System.Loaded = true
	return &System, nil
}

func Auth(ac string) error {
	if ac != System.API.AuthKey {
		msg := "Invalid auth code."
		log.Warning(msg)
		return errors.New(msg)
	}
	return nil
}

// ==============================================================================================================================
//                                      SYSTEM
// ==============================================================================================================================
// ------------------------------- SystemType -------------------------------
type SystemType struct {
	Loaded          bool
	Instrumentation DebugType `json:"instrumentation"`
	API             API_Type  `json:"api"`
	Data            Data_Type `json:"data"`
}

func (x *SystemType) Display() string {
	s := fmt.Sprintf("\n==================================== SYSTEM ==================================\n")
	s += spew.Sdump(x)
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

// ==============================================================================================================================
//                                      REPORTS
// ==============================================================================================================================

// ------------------------------- ReportsType -------------------------------
type Data_Type struct {
	lastId  int64
	Reports []*Report_Type `json:"reports"`
	sync.Mutex
}

func (d *Data_Type) FindDeviceId(id string) ([]*Report_Type, error) {
	out := make([]*Report_Type, 0)
	for _, v := range d.Reports {
		if v.DeviceId == id {
			out = append(out, v)
		}
	}
	return out, nil
}

// ------------------------------- TableType -------------------------------
type Report_Type struct {
	Id                int64             `json:"id"`
	DateCreated       common.CustomTime `json:"datecreated"`
	DeviceType        string            `json:"devicetype"`
	DeviceModel       string            `json:"devicemodel"`
	DeviceId          string            `json:"deviceid"`
	RequestType       string            `json:"requesttype"`
	RequestTypeId     string            `json:"requesttypeid"`
	Latitude          float64           `json:"latitude"`
	Longitude         float64           `json:"longitude"`
	Directionality    string            `json:"directionality"`
	Description       string            `json:"description"`
	AuthorNameFirst   string            `json:"authornamefirst"`
	AuthorNameLast    string            `json:"authornamelast"`
	AuthorEmail       string            `json:"authoremail"`
	AuthorTelephone   string            `json:"authortelephone"`
	AuthorIsAnonymous bool              `json:"authorisanonymous"`
}

func (r *Report_Type) Distance(rlat, rlon float64) float64 {
	return Distance(rlat, rlon, r.Latitude, r.Longitude)
}

// Displays the contents of the Spec_Type custom type.
func (s Report_Type) String() string {
	ls := new(logs.LogString)
	ls.AddS("Report\n")
	ls.AddF("Id: %v\n", s.Id)
	ls.AddF("DateCreated \"%v\"\n", s.DateCreated)
	ls.AddF("Device - type %s  model: %s  Id: %s\n", s.DeviceType, s.DeviceModel, s.DeviceId)
	ls.AddF("Request - type: %q  id: %q\n", s.RequestType, s.RequestTypeId)
	ls.AddF("Location - lat: %v  lon: %v  directionality: %q\n", s.Latitude, s.Longitude, s.Directionality)
	ls.AddF("Description: %q\n", s.Description)
	ls.AddF("Author(anon: %t) %s %s  Email: %s  Tel: %s\n", s.AuthorIsAnonymous, s.AuthorNameFirst, s.AuthorNameLast, s.AuthorEmail, s.AuthorTelephone)
	return ls.Box(90)
}
