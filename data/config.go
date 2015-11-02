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
	Config ConfigType
	Data   Data_Type
)

func ReadConfig(filePathConfig, filePathData string) error {
	log.Info("Loading configuration file - Config: %q", filePathConfig)
	_, errConfig := readConfig(filePathConfig)
	if errConfig != nil {
		return errors.New(fmt.Sprintf("Error loading config: %s", errConfig))
	}

	_, errData := readData(filePathData)
	if errData != nil {
		return errors.New(fmt.Sprintf("Error loading config: %s", errConfig))
	}

	return nil
}

func Auth(ac string) bool {
	if ac == Config.API.AuthKey {
		return true
	}
	return false
}

// ==============================================================================================================================
//                                      Config
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
	if Config.Loaded {
		msg := "Duplicate calls to load Config file!"
		log.Warning(msg)
		return &Config, errors.New(msg)
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		msg := fmt.Sprintf("Unable to open the Config Config file - specified at: %q.\nError: %v", filePath, err)
		log.Critical(msg)
		return nil, errors.New(msg)
	}

	err = json.Unmarshal([]byte(file), &Config)
	if err != nil {
		msg := fmt.Sprintf("Invalid JSON in the Config Config file - specified at: %q.\nError: %v", filePath, err)
		log.Critical(msg)
		return nil, errors.New(msg)
	}

	fmt.Printf("After loading dates...\n%s\n", spew.Sdump(Config))

	Config.Loaded = true
	return &Config, nil
}

// ==============================================================================================================================
//                                      DATA
// ==============================================================================================================================
// ------------------------------- Data_Type -------------------------------
type Data_Type struct {
	Loaded bool
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

func (x *Data_Type) Display() string {
	s := fmt.Sprintf("\n==================================== DATA ==================================\n")
	s += spew.Sdump(x)
	return s
}


func readData(filePath string) (*Data_Type, error) {
	if Data.Loaded {
		msg := "Duplicate calls to load Data file!"
		log.Warning(msg)
		return &Data, errors.New(msg)
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		msg := fmt.Sprintf("Unable to open the Data Data file - specified at: %q.\nError: %v", filePath, err)
		log.Critical(msg)
		return nil, errors.New(msg)
	}

	err = json.Unmarshal([]byte(file), &Data)
	if err != nil {
		msg := fmt.Sprintf("Invalid JSON in the Data Data file - specified at: %q.\nError: %v", filePath, err)
		log.Critical(msg)
		return nil, errors.New(msg)
	}

	// Update Last ID
	var lastId int64 = 0
	for _, v := range Data.Reports {
		if v.Id > lastId {
			lastId = v.Id
		}
	}
	Data.lastId = lastId

	fmt.Printf("After loading dates...\n%s\n", spew.Sdump(Data))

	Data.Loaded = true
	return &Data, nil
}

// ==============================================================================================================================
//                                      REPORTS
// ==============================================================================================================================

// ------------------------------- ReportsType -------------------------------


// ------------------------------- Report_Type -------------------------------
type Report_Type struct {
	Id                int64             `json:"Id" xml:"Id"`
	DateCreated       common.CustomTime `json:"DateCreated" xml:"DateCreated"`
	DateUpdated       common.CustomTime `json:"DateUpdated" xml:"DateUpdated"`
	DeviceType        string            `json:"DeviceType" xml:"DeviceType"`
	DeviceModel       string            `json:"DeviceModel" xml:"DeviceModel"`
	DeviceId          string            `json:"DeviceId" xml:"DeviceId"`
	RequestType       string            `json:"RequestType" xml:"RequestType"`
	RequestTypeId     string            `json:"RequestTypeId" xml:"RequestTypeId"`
	ImageUrl          string            `json:"ImageUrl" xml:"ImageUrl"`
	ImageUrlXl        string            `json:"ImageUrlXl" xml:"ImageUrlXl"`
	ImageUrlLg        string            `json:"ImageUrlLg" xml:"ImageUrlLg"`
	ImageUrlMd        string            `json:"ImageUrlMd" xml:"ImageUrlMd"`
	ImageUrlSm        string            `json:"ImageUrlSm" xml:"ImageUrlSm"`
	ImageUrlXs        string            `json:"ImageUrlXs" xml:"ImageUrlXs"`
	City              string            `json:"City" xml:"City"`
	State             string            `json:"State" xml:"State"`
	ZipCode           string            `json:"ZipCode" xml:"ZipCode"`
	Latitude          float64           `json:"Latitude" xml:"Latitude"`
	Longitude         float64           `json:"Longitude" xml:"Longitude"`
	Directionality    string            `json:"Directionality" xml:"Directionality"`
	Description       string            `json:"Description" xml:"Description"`
	AuthorNameFirst   string            `json:"AuthorNameFirst" xml:"AuthorNameFirst"`
	AuthorNameLast    string            `json:"AuthorNameLast" xml:"AuthorNameLast"`
	AuthorEmail       string            `json:"AuthorEmail" xml:"AuthorEmail"`
	AuthorTelephone   string            `json:"AuthorTelephone" xml:"AuthorTelephone"`
	AuthorIsAnonymous bool              `json:"AuthorIsAnonymous" xml:"AuthorIsAnonymous"`
	UrlDetail         string            `json:"UrlDetail" xml:"UrlDetail"`
	UrlShortened      string            `json:"UrlShortened" xml:"UrlShortened"`
	StatusType        string            `json:"StatusType" xml:"StatusType"`
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
	ls.AddF("          %s, %s   %s\n", s.City, s.State, s.ZipCode)
	ls.AddF("Description: %q\n", s.Description)
	ls.AddF("Images - std: %s\n", s.ImageUrl)
	ls.AddF("          XS: %s\n", s.ImageUrlXs)
	ls.AddF("          SM: %s\n", s.ImageUrlSm)
	ls.AddF("          MD: %s\n", s.ImageUrlMd)
	ls.AddF("          LG: %s\n", s.ImageUrlLg)
	ls.AddF("          XL: %s\n", s.ImageUrlXl)
	ls.AddF("Author(anon: %t) %s %s  Email: %s  Tel: %s\n", s.AuthorIsAnonymous, s.AuthorNameFirst, s.AuthorNameLast, s.AuthorEmail, s.AuthorTelephone)
	return ls.Box(90)
}
