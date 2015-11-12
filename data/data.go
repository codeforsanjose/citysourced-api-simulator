package data

import (
	"CitySourcedAPI/geo"
	"CitySourcedAPI/logs"

	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/davecgh/go-spew/spew"
)

var (
	log = logs.Log
	D   Reports
)

func Init(fileName string) error {
	D.Reports = NewReportList()
	_, err := readData(fileName)
	if err != nil {
		return fmt.Errorf("Error loading config: %s", err)
	}
	return nil
}

// ==============================================================================================================================
//                                      DATA
// ==============================================================================================================================

// ------------------------------- Data_Type -------------------------------
type Reports struct {
	Loaded  bool
	lastID  int64
	Reports ReportList `json:"reports" xml:"reports"`
	// Reports []*Report_Type `json:"reports" xml:"reports"`
	indID map[int64]*Report
	sync.Mutex
}

func (d *Reports) LastID() int64 {
	return d.lastID
}

func (d *Reports) Validate() error {
	for _, r := range d.Reports {
		r.Validate()
	}
	return nil
}

func (d *Reports) Append(st BaseReport) error {
	if err := st.Validate(); err != nil {
		return err
	}
	// log.Debug("[AddReport] st: type: %T\n%s", st, spew.Sdump(st))
	d.Lock()
	d.lastID = d.lastID + 1
	r, _ := d.Reports.AddBR(d.lastID, &st)
	d.indID[d.lastID] = r
	d.Unlock()
	log.Debug(d.Display())
	return nil
}

func (d *Reports) FindDeviceID(id string) ([]*Report, error) {
	rlist := NewReportList()
	for _, v := range d.Reports {
		if v.DeviceID == id {
			rlist = append(rlist, v)
		}
	}
	return rlist, nil
}

func (d *Reports) FindID(id int64) (*Report, error) {
	r := d.indID[id]
	if r != nil {
		return r, nil
	}
	return r, errors.New(fmt.Sprintf("ID: %v not found!", id))
}

func (d *Reports) FindAddress(addr string, radius float64, limit int64) ([]*Report, error) {
	rlist := NewReportListD()
	log.Debug("FindAddress - addr: %s  radius: %v", addr, radius)
	lat, lng, e := geo.GetLatLng(addr)
	if e != nil {
		msg := fmt.Sprintf("GeoLoc failed for address: %s", e)
		log.Warning(msg)
		return rlist.ReportList, errors.New(msg)
	}
	return d.FindLL(lat, lng, radius, limit)
}

func (d *Reports) FindLL(lat, lng, radius float64, limit int64) ([]*Report, error) {
	rlist := NewReportListD()
	log.Debug("Scanning Reports for reports within %v meters of: %v|%v", radius, lat, lng)
	for _, v := range d.Reports {
		dist := Distance(lat, lng, v.LatitudeV, v.LongitudeV)
		fmt.Printf("ID: %v  dist: %v\n", v.ID, dist)
		if dist < radius {
			rlist.Add(v, dist)
		}
	}

	rlist.Sort()
	rlist.Limit(limit)

	log.Debug(">>> rlist:\n%s\n", spew.Sdump(rlist))
	return rlist.ReportList, nil
}

func (x *Reports) Display() string {
	s := fmt.Sprintf("\n==================================== DATA ==================================\n")
	s += spew.Sdump(x)
	return s
}

func readData(filePath string) (*Reports, error) {
	if D.Loaded {
		msg := "Duplicate calls to load Data file!"
		log.Warning(msg)
		return &D, errors.New(msg)
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		msg := fmt.Sprintf("Failed to %s", err)
		log.Critical(msg)
		return nil, errors.New(msg)
	}

	err = json.Unmarshal([]byte(file), &D)
	if err != nil {
		msg := fmt.Sprintf("Invalid JSON in the Data file %q: %s", filePath, err)
		log.Critical(msg)
		return nil, errors.New(msg)
	}

	err = D.Validate()
	if err != nil {
		msg := fmt.Sprintf("Unable to validate data (check lng, lat, etc): %q: %s", filePath, err)
		log.Critical(msg)
		return nil, errors.New(msg)
	}

	// Build Indexes
	D.index()
	fmt.Println(spew.Sdump(D.indID))

	// Update Last ID
	var lastID int64 = 0
	for _, v := range D.Reports {
		if v.ID > lastID {
			lastID = v.ID
		}
	}
	D.lastID = lastID

	D.Loaded = true
	return &D, nil
}

func (d *Reports) index() error {
	d.indID = make(map[int64]*Report)
	for _, r := range d.Reports {
		d.indID[r.ID] = r
	}
	return nil
}
