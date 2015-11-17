package request

import (
	"CitySourcedAPI/data"
	"CitySourcedAPI/logs"
	"CitySourcedAPI/response"

	"time"

	"github.com/davecgh/go-spew/spew"
)

// ==============================================================================================================================
//                                      GetReportsByAddress
// ==============================================================================================================================

type GetReportsByAddress struct {
	Request
	Processor
	Address        string          `xml:"Address" json:"Address"`
	Radius         string          `xml:"Radius" json:"Radius"`
	radius         float64         //
	MaxResults     string          `xml:"MaxResults" json:"MaxResults"`
	maxResults     int64           //
	IncludeDetails string          `xml:"IncludeDetails" json:"IncludeDetails"`
	includeDetails bool            //
	DateRangeStart data.CustomTime `xml:"DateRangeStart" json:"DateRangeStart"`
	DateRangeEnd   data.CustomTime `xml:"DateRangeEnd" json:"DateRangeEnd"`
	CurrentStatus  string          `xml:"CurrentStatus" json:"CurrentStatus"`
}

func (st *GetReportsByAddress) Validate(start time.Time) string {
	var v validate
	st.start = start
	st.radius = v.float("Radius", st.Radius)
	st.maxResults = v.int("MaxResults", st.MaxResults)
	st.includeDetails = v.bool("IncludeDetails", st.IncludeDetails)
	return v.errmsg
}

func (st *GetReportsByAddress) Run() (string, error) {
	rpts, _ := data.FindAddress(st.Address, st.radius, st.maxResults)
	log.Debug(">>> rpts:\n%s\n", spew.Sdump(rpts))

	resp, _ := response.NewResponseReports(true, st.Start(), rpts)
	return resp, nil
}

func (st GetReportsByAddress) String() string {
	ls := new(logs.LogString)
	ls.AddS("GetReportsByAddress\n")
	ls.AddS(st.Request.String())
	ls.AddF("Address \"%v\"\n", st.Address)
	ls.AddF("Radius %s/%v   MaxResults: %s/%v\n", st.Radius, st.radius, st.MaxResults, st.maxResults)
	ls.AddF("IncludeDetails: %v/%t\n", st.IncludeDetails, st.includeDetails)
	ls.AddF("Date Range: %v  to: %v \n", st.DateRangeStart, st.DateRangeEnd)
	return ls.Box(90)
}

// ==============================================================================================================================
//                                      GetReportsByLatLng
// ==============================================================================================================================

type GetReportsByLatLng struct {
	Request
	Processor
	Latitude       string          `xml:"Latitude" json:"Latitude"`
	LatitudeV      float64         //
	Longitude      string          `xml:"Longitude" json:"Longitude"`
	LongitudeV     float64         //
	Radius         string          `xml:"Radius" json:"Radius"`
	radius         float64         //
	MaxResults     string          `xml:"MaxResults" json:"MaxResults"`
	maxResults     int64           //
	IncludeDetails string          `xml:"IncludeDetails" json:"IncludeDetails"`
	includeDetails bool            //
	DateRangeStart data.CustomTime `xml:"DateRangeStart" json:"DateRangeStart"`
	DateRangeEnd   data.CustomTime `xml:"DateRangeEnd" json:"DateRangeEnd"`
	CurrentStatus  string          `xml:"CurrentStatus" json:"CurrentStatus"`
}

func (st *GetReportsByLatLng) Validate(start time.Time) string {
	var v validate
	st.start = start
	st.LatitudeV = v.float("Latitude", st.Latitude)
	st.LongitudeV = v.float("Longitude", st.Longitude)
	st.radius = v.float("Radius", st.Radius)
	st.maxResults = v.int("MaxResults", st.MaxResults)
	st.includeDetails = v.bool("IncludeDetails", st.IncludeDetails)
	return v.errmsg
}

func (st *GetReportsByLatLng) Run() (string, error) {
	rpts, _ := data.FindLL(st.LatitudeV, st.LongitudeV, st.radius, st.maxResults)
	log.Debug(">>> rpts:\n%s\n", spew.Sdump(rpts))

	resp, _ := response.NewResponseReports(true, st.Start(), rpts)
	return resp, nil
}

func (st GetReportsByLatLng) String() string {
	ls := new(logs.LogString)
	ls.AddS("GetReportsByLatLng\n")
	ls.AddS(st.Request.String())
	ls.AddF("Loc \"%v:%v\"\n", st.LatitudeV, st.LongitudeV)
	ls.AddF("Radius %s/%v\n", st.Radius, st.radius)
	ls.AddF("MaxResults: %s/%v\n", st.MaxResults, st.maxResults)
	ls.AddF("IncludeDetails: %v/%t\n", st.IncludeDetails, st.includeDetails)
	ls.AddF("Date Range: %v  to: %v \n", st.DateRangeStart, st.DateRangeEnd)
	return ls.Box(90)
}

// ==============================================================================================================================
//                                      GetReportsByDeviceId
// ==============================================================================================================================

type GetReportsByDeviceID struct {
	Request
	Processor
	DeviceID       string          `xml:"DeviceId" json:"DeviceId"`
	MaxResults     string          `xml:"MaxResults" json:"MaxResults"`
	maxResults     int64           //
	IncludeDetails string          `xml:"IncludeDetails" json:"IncludeDetails"`
	includeDetails bool            //
	DateRangeStart data.CustomTime `xml:"DateRangeStart" json:"DateRangeStart"`
	DateRangeEnd   data.CustomTime `xml:"DateRangeEnd" json:"DateRangeEnd"`
	CurrentStatus  string          `xml:"CurrentStatus" json:"CurrentStatus"`
}

func (st *GetReportsByDeviceID) Validate(start time.Time) string {
	var v validate
	st.start = start
	st.maxResults = v.int("MaxResults", st.MaxResults)
	st.includeDetails = v.bool("IncludeDetails", st.IncludeDetails)
	return v.errmsg
}

func (st *GetReportsByDeviceID) Run() (string, error) {
	rpts, _ := data.FindDeviceID(st.DeviceID)
	log.Debug(">>> rpts:\n%s\n", spew.Sdump(rpts))

	resp, _ := response.NewResponseReports(true, st.Start(), rpts)
	return resp, nil
}

func (st GetReportsByDeviceID) String() string {
	ls := new(logs.LogString)
	ls.AddS("GetReportsByDeviceID\n")
	ls.AddS(st.Request.String())
	ls.AddF("DeviceID: %q\n", st.DeviceID)
	ls.AddF("MaxResults: %s/%v\n", st.MaxResults, st.maxResults)
	ls.AddF("IncludeDetails: %v/%t\n", st.IncludeDetails, st.includeDetails)
	ls.AddF("Date Range: %v  to: %v \n", st.DateRangeStart, st.DateRangeEnd)
	return ls.Box(90)
}

// ==============================================================================================================================
//                                      GetReportsByZipCode
// ==============================================================================================================================

type GetReportsByZipCode struct {
	Request
	Processor
	ZipCode        string          `xml:"ZipCode" json:"ZipCode"`
	MaxResults     string          `xml:"MaxResults" json:"MaxResults"`
	maxResults     int64           //
	IncludeDetails string          `xml:"IncludeDetails" json:"IncludeDetails"`
	includeDetails bool            //
	DateRangeStart data.CustomTime `xml:"DateRangeStart" json:"DateRangeStart"`
	DateRangeEnd   data.CustomTime `xml:"DateRangeEnd" json:"DateRangeEnd"`
	CurrentStatus  string          `xml:"CurrentStatus" json:"CurrentStatus"`
}

func (st *GetReportsByZipCode) Validate(start time.Time) string {
	var v validate
	st.start = start
	st.maxResults = v.int("MaxResults", st.MaxResults)
	st.includeDetails = v.bool("IncludeDetails", st.IncludeDetails)
	return v.errmsg
}

func (st *GetReportsByZipCode) Run() (string, error) {
	rpts, _ := data.FindZipCode(st.ZipCode)
	log.Debug(">>> rpts:\n%s\n", spew.Sdump(rpts))

	resp, _ := response.NewResponseReports(true, st.Start(), rpts)
	return resp, nil
}

func (st GetReportsByZipCode) String() string {
	ls := new(logs.LogString)
	ls.AddS("GetReportsByZipCode\n")
	ls.AddS(st.Request.String())
	ls.AddF("ZipCode: %q\n", st.ZipCode)
	ls.AddF("MaxResults: %s/%v\n", st.MaxResults, st.maxResults)
	ls.AddF("IncludeDetails: %v/%t\n", st.IncludeDetails, st.includeDetails)
	ls.AddF("Date Range: %v  to: %v \n", st.DateRangeStart, st.DateRangeEnd)
	return ls.Box(90)
}

// ==============================================================================================================================
//                                      GetReport
// ==============================================================================================================================

type GetReport struct {
	Request
	Processor
	ReportID        string          `xml:"ReportId" json:"ReportId"`
	ReportIDV       int64           //
	IncludeDetails  string          `xml:"IncludeDetails" json:"IncludeDetails"`
	includeDetails  bool            //
	IncludeComments string          `xml:"IncludeComments" json:"IncludeComments"`
	includeComments bool            //
	IncludeVotes    string          `xml:"IncludeVotes" json:"IncludeVotes"`
	includeVotes    bool            //
	DateRangeStart  data.CustomTime `xml:"DateRangeStart" json:"DateRangeStart"`
	DateRangeEnd    data.CustomTime `xml:"DateRangeEnd" json:"DateRangeEnd"`
	CurrentStatus   string          `xml:"CurrentStatus" json:"CurrentStatus"`
}

func (st *GetReport) Validate(start time.Time) string {
	var v validate
	st.start = start
	st.ReportIDV = v.int("ReportID", st.ReportID)
	st.includeDetails = v.bool("IncludeDetails", st.IncludeDetails)
	st.includeComments = v.bool("IncludeComments", st.IncludeComments)
	st.includeVotes = v.bool("IncludeVotes", st.IncludeVotes)
	return v.errmsg
}

func (st *GetReport) Run() (string, error) {
	rpts, _ := data.FindID(st.ReportIDV)
	log.Debug(">>> rpts:\n%s\n", spew.Sdump(rpts))

	resp, _ := response.NewResponseReports(true, st.Start(), rpts)
	return resp, nil
}

func (st GetReport) String() string {
	ls := new(logs.LogString)
	ls.AddS("GetReport\n")
	ls.AddS(st.Request.String())
	ls.AddF("ReportID: %q\n", st.ReportID)
	ls.AddF("IncludeComments: %s/%v\n", st.IncludeComments, st.includeComments)
	ls.AddF("IncludeDetails: %v/%t\n", st.IncludeDetails, st.includeDetails)
	ls.AddF("IncludeVotes: %v/%t\n", st.IncludeVotes, st.includeVotes)
	ls.AddF("Date Range: %v  to: %v \n", st.DateRangeStart, st.DateRangeEnd)
	return ls.Box(90)
}
