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
	Address        string `xml:"Address" json:"Address"`
	Radius         string `xml:"Radius" json:"Radius"`
	radius         float64
	MaxResults     string `xml:"MaxResults" json:"MaxResults"`
	maxResults     int64
	IncludeDetails string `xml:"IncludeDetails" json:"IncludeDetails"`
	includeDetails bool
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
	rpts, _ := data.D.FindAddress(st.Address, st.radius, st.maxResults)
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
	Latitude       string  `xml:"Latitude" json:"Latitude"`
	LatitudeV      float64 //
	Longitude      string  `xml:"Longitude" json:"Longitude"`
	LongitudeV     float64 //
	Radius         string  `xml:"Radius" json:"Radius"`
	radius         float64
	MaxResults     string `xml:"MaxResults" json:"MaxResults"`
	maxResults     int64
	IncludeDetails string `xml:"IncludeDetails" json:"IncludeDetails"`
	includeDetails bool
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
	rpts, _ := data.D.FindLL(st.LatitudeV, st.LongitudeV, st.radius, st.maxResults)
	log.Debug(">>> rpts:\n%s\n", spew.Sdump(rpts))

	resp, _ := response.NewResponseReports(true, st.Start(), rpts)

	return resp, nil
}

func (st GetReportsByLatLng) String() string {
	ls := new(logs.LogString)
	ls.AddS("GetReportsByLatLng\n")
	ls.AddS(st.Request.String())
	ls.AddF("Loc \"%v:%v\"\n", st.LatitudeV, st.LongitudeV)
	ls.AddF("Radius %s/%v   MaxResults: %s/%v\n", st.Radius, st.radius, st.MaxResults, st.maxResults)
	ls.AddF("IncludeDetails: %v/%t\n", st.IncludeDetails, st.includeDetails)
	ls.AddF("Date Range: %v  to: %v \n", st.DateRangeStart, st.DateRangeEnd)
	return ls.Box(90)
}
