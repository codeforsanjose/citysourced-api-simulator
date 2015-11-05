package data

import (
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
	_, err := readData(fileName)
	if err != nil {
		return errors.New(fmt.Sprintf("Error loading config: %s", err))
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
	Reports []*Report_Type `json:"reports"`
	indId   map[int64]*Report_Type
	sync.Mutex
}

func (d *Data_Type) LastId() int64 {
	return d.lastId
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

func (d *Data_Type) FindId(id int64) (*Report_Type, error) {
	r := d.indId[id]
	if r != nil {
		return r, nil
	}
	return r, errors.New(fmt.Sprintf("Id: %v not found!", id))
}

func (d *Data_Type) FindAddress(addr string, radius float64) ([]*Report_Type, error) {
	rlist := make([]*Report_Type, 0)
	log.Debug("FindAddress - addr: %s  radius: %v", addr, radius)
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
