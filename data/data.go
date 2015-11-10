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
	D   Data_Type
)

func Init(fileName string) error {
	D.Reports = make([]*Report_Type, 0)
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
type Data_Type struct {
	Loaded  bool
	lastId  int64
	Reports ReportList `json:"reports" xml:"reports"`
	// Reports []*Report_Type `json:"reports" xml:"reports"`
	indId map[int64]*Report_Type
	sync.Mutex
}

func (d *Data_Type) LastId() int64 {
	return d.lastId
}

func (d *Data_Type) AddReport(st BaseReport) error {
	if err := st.Validate(); err != nil {
		return err
	}
	// log.Debug("[AddReport] st: type: %T\n%s", st, spew.Sdump(st))
	d.Lock()
	d.lastId = d.lastId + 1
	r, _ := d.Reports.AddBR(d.lastId, &st)
	d.indId[d.lastId] = r
	d.Unlock()
	log.Debug(d.Display())
	return nil
}

func (d *Data_Type) FindDeviceId(id string) ([]*Report_Type, error) {
	rlist := make([]*Report_Type, 0)
	for _, v := range d.Reports {
		if v.DeviceId == id {
			rlist = append(rlist, v)
		}
	}
	return rlist, nil
}

func (d *Data_Type) FindId(id int64) (*Report_Type, error) {
	r := d.indId[id]
	if r != nil {
		return r, nil
	}
	return r, errors.New(fmt.Sprintf("Id: %v not found!", id))
}

func (d *Data_Type) Validate() error {
	for _, r := range d.Reports {
		r.Validate()
	}
	return nil
}

func (d *Data_Type) FindAddress(addr string, radius float64) ([]*Report_Type, error) {
	rlist := NewReportList()
	log.Debug("FindAddress - addr: %s  radius: %v", addr, radius)
	alat, alng, e := geo.GetLatLng(addr)
	if e != nil {
		msg := fmt.Sprintf("GeoLoc failed for address: %s", e)
		log.Warning(msg)
		return rlist, errors.New(msg)
	}
	log.Debug("Scanning Reports for reports within %v meters of: %v|%v", radius, alat, alng)
	for _, v := range d.Reports {
		dist := Distance(alat, alng, v.latitude, v.longitude)
		fmt.Printf("Id: %v  dist: %v\n", v.Id, dist)
		if dist < radius {
			rlist.Add(v)
		}
	}
	log.Debug(">>> rlist:\n%s\n", spew.Sdump(rlist))
	return rlist, nil
}

func (x *Data_Type) Display() string {
	s := fmt.Sprintf("\n==================================== DATA ==================================\n")
	s += spew.Sdump(x)
	return s
}

func readData(filePath string) (*Data_Type, error) {
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
	fmt.Println(spew.Sdump(D.indId))

	// Update Last ID
	var lastId int64 = 0
	for _, v := range D.Reports {
		if v.Id > lastId {
			lastId = v.Id
		}
	}
	D.lastId = lastId

	D.Loaded = true
	return &D, nil
}

func (d *Data_Type) index() error {
	d.indId = make(map[int64]*Report_Type)
	for _, r := range d.Reports {
		d.indId[r.Id] = r
	}
	return nil
}
