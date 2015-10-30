package data

import (
	"CitySourcedAPI/logs"
	"encoding/json"

	"github.com/davecgh/go-spew/spew"

	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

var debug = true
var verbose = true
var log = logs.Log

// ==============================================================================================================================
//                                      CONFIG VARIABLES
// ==============================================================================================================================

var System SystemType

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

	for _, rpt := range System.Data.Reports {
		rpt.parseDate()
	}

	fmt.Printf("After loading dates...\n%s\n", spew.Sdump(System.Data.Reports))

	System.Loaded = true
	return &System, nil
}

// ==============================================================================================================================
//                                      SYSTEM
// ==============================================================================================================================
// ------------------------------- SystemType -------------------------------
type SystemType struct {
	Loaded          bool
	Instrumentation DebugType `json:"instrumentation"`
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

// ==============================================================================================================================
//                                      REPORTS
// ==============================================================================================================================

// ------------------------------- ReportsType -------------------------------
type Data_Type struct {
	Reports []*Report_Type `json:"reports"`
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
	Id                int64  `json:"id"`
	DateCreated       string `json:"datecreated"`
	DC                time.Time
	DeviceType        string  `json:"devicetype"`
	DeviceModel       string  `json:"devicemodel"`
	DeviceId          string  `json:"deviceid"`
	RequestType       string  `json:"requesttype"`
	RequestTypeId     string  `json:"requesttypeid"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	Directionality    string  `json:"directionality"`
	Description       string  `json:"description"`
	AuthorNameFirst   string  `json:"authornamefirst"`
	AuthorNameLast    string  `json:"authornamelast"`
	AuthorEmail       string  `json:"authoremail"`
	AuthorTelephone   string  `json:"authortelephone"`
	AuthorIsAnonymous bool    `json:"authorisanonymous"`
}

func (r *Report_Type) parseDate() {
	dt, _ := time.Parse("2006-01-02T15:04:05", r.DateCreated)
	r.DC = dt
	fmt.Printf("*** DateCreated: %s  dt: %v (%T)   %v\n", r.DateCreated, dt, dt, r.DC)
}

func (r *Report_Type) Distance(rlat, rlon float64) float64 {
	return Distance(rlat, rlon, r.Latitude, r.Longitude)
}

// func (x *Report_Type) String() string {
// 	return spew.Sdump(x)
// }


// ==============================================================================================================================
//                                      Custom Time Format
// ==============================================================================================================================
type CustomTime struct {
    time.Time
}

const ctLayout = "1974-05-20T13:45:30"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
    ct.Time, err = time.Parse(ctLayout, string(b))
    return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
    return []byte(ct.Time.Format(ctLayout)), nil
}

var nilTime = (time.Time{}).UnixNano()
func (ct *CustomTime) IsSet() bool {
    return ct.UnixNano() != nilTime
}

