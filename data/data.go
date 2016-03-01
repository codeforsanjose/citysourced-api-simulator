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
	log     = logs.Log
	rptData Reports
)

func Init(fileName string) error {
	log.Info("Loading data file: %q", fileName)

	// Reports
	rptData.Reports = newReportList(false)
	_, err := readReportData(fileName)
	if err != nil {
		return fmt.Errorf("Error loading Report Data: %s", err)
	}

	// Comments
	_, err = readCommentData(fileName)
	if err != nil {
		return fmt.Errorf("Error loading Comment Data: %s", err)
	}

	return nil
}

func LastID() int64 {
	return rptData.lastID
}

func Append(st BaseReport, comments bool) (int64, int64, error) {
	if err := st.Validate(); err != nil {
		return 0, 0, err
	}
	// log.Debug("[AddReport] st: type: %T\n%s", st, spew.Sdump(st))
	rptData.Lock()
	rptData.lastID = rptData.lastID + 1
	r, _ := rptData.Reports.AddBR(rptData.lastID, &st, comments)
	rptData.indID[rptData.lastID] = r
	rptData.Unlock()
	log.Debug(DisplayReports())
	return rptData.lastID, 99, nil
}

func GetID(id int64, comments bool) (*Report, error) {
	r := rptData.indID[id]
	if r != nil {
		return r, nil
	}
	return r, fmt.Errorf("ID: %v not found", id)
}

func FindID(id int64, comments bool) ([]*Report, error) {
	rlist := newReportList(false)
	if r, ok := rptData.indID[id]; ok {
		log.Debug(spew.Sdump(r))
		if comments {
			c, err := FindReportComments(id)
			if err == nil {
				r.Comments = c
			}
		}
		log.Debug(spew.Sdump(r))
		rlist = append(rlist, r)
	}
	return rlist, nil
}

func FindDeviceID(id string, comments bool) ([]*Report, error) {
	rlist := newReportList(false)
	for _, v := range rptData.Reports {
		if v.DeviceID == id {
			rlist = append(rlist, v)
		}
	}
	return rlist, nil
}

func FindZipCode(zip string, comments bool) ([]*Report, error) {
	rlist := newReportList(false)
	for _, v := range rptData.Reports {
		if v.ZipCode == zip {
			rlist = append(rlist, v)
		}
	}
	return rlist, nil
}

func FindAddress(addr string, radius float64, limit int64, comments bool) ([]*Report, error) {
	rlist := newReportListD(false)
	log.Debug("FindAddress - addr: %s  radius: %v", addr, radius)
	lat, lng, e := geo.GetLatLng(addr)
	if e != nil {
		msg := fmt.Sprintf("GeoLoc failed for address: %s", e)
		log.Warning(msg)
		return rlist.ReportList, errors.New(msg)
	}
	return FindLL(lat, lng, radius, limit, comments)
}

func FindLL(lat, lng, radius float64, limit int64, comments bool) ([]*Report, error) {
	rlist := newReportListD(false)
	log.Debug("Scanning Reports for reports within %v meters of: %v|%v", radius, lat, lng)
	for _, v := range rptData.Reports {
		dist := Distance(lat, lng, v.LatitudeV, v.LongitudeV)
		log.Debug("ID: %v  dist: %v\n", v.ID, dist)
		if dist < radius {
			rlist.Add(v, dist, comments)
		}
	}

	rlist.Sort()
	rlist.Limit(limit)

	log.Debug(">>> rlist:\n%s\n", spew.Sdump(rlist))
	return rlist.ReportList, nil
}

func UpdateSLA(id int64, sla string) error {
	return rptData.updateSLA(id, sla)
}

func Vote(id int64) error {
	return rptData.vote(id)
}

func DisplayReports() string {
	s := fmt.Sprintf("\n==================================== DATA ==================================\n")
	s += spew.Sdump(rptData)
	return s
}

// This is for running "go test" only.  It should be commented out after testing.
func ReportDataTEST() *ReportList {
	return &rptData.Reports
}

// ==============================================================================================================================
//                                      DATA
// ==============================================================================================================================

// ------------------------------- Data_Type -------------------------------
type Reports struct {
	Loaded  bool
	lastID  int64
	Reports ReportList `json:"reports" xml:"reports"`
	indID   map[int64]*Report
	sync.Mutex
}

func (d *Reports) validate() error {
	for _, r := range d.Reports {
		r.Validate()
	}
	return nil
}

func (d *Reports) index() error {
	d.indID = make(map[int64]*Report)
	for _, r := range d.Reports {
		d.indID[r.ID] = r
	}
	return nil
}

func (d *Reports) updateSLA(id int64, sla string) error {
	rpt, ok := d.indID[id]
	if !ok {
		msg := fmt.Sprintf("ID: %d does not exist", id)
		log.Warning(msg)
		return errors.New(msg)
	}
	rpt.updateSLA(sla)
	return nil
}

func (d *Reports) vote(id int64) error {
	rpt, ok := d.indID[id]
	if !ok {
		msg := fmt.Sprintf("ID: %d does not exist", id)
		log.Warning(msg)
		return errors.New(msg)
	}
	rpt.vote()
	return nil
}

func readReportData(filePath string) (*Reports, error) {
	if rptData.Loaded {
		msg := "Duplicate calls to load Report Data file!"
		log.Warning(msg)
		return &rptData, errors.New(msg)
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		msg := fmt.Sprintf("Failed to %s", err)
		log.Critical(msg)
		return nil, errors.New(msg)
	}

	err = json.Unmarshal([]byte(file), &rptData)
	if err != nil {
		msg := fmt.Sprintf("Invalid JSON in the Data file %q: %s", filePath, err)
		log.Critical(msg)
		return nil, errors.New(msg)
	}

	err = rptData.validate()
	if err != nil {
		msg := fmt.Sprintf("Unable to validate data (check lng, lat, etc): %q: %s", filePath, err)
		log.Critical(msg)
		return nil, errors.New(msg)
	}

	// Build Indexes
	rptData.index()
	log.Debug(spew.Sdump(rptData.indID))

	// Update Last ID
	var lastID int64
	for _, v := range rptData.Reports {
		if v.ID > lastID {
			lastID = v.ID
		}
	}
	rptData.lastID = lastID

	rptData.Loaded = true
	return &rptData, nil
}
