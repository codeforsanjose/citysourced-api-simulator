package request

import (
	"CitySourcedAPI/data"
	"CitySourcedAPI/logs"
	"CitySourcedAPI/response"

	"encoding/xml"
	"errors"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// ==============================================================================================================================
//                                      GetReportsByAddress
// ==============================================================================================================================

func GetReportsByAddress(input string, start time.Time) (string, error) {
	st := new(GetReportsByAddressType)
	err := xml.Unmarshal([]byte(input), st)
	if err != nil {
		msg := fmt.Sprintf("Unable to unmarshal GetReportsByAddress request: %s", err)
		log.Warning(msg)
		return "", errors.New(msg)
	}

	st.start = start
	if errmsg := st.validate(); errmsg != "" {
		msg := fmt.Sprintf("Invalid request - %s", errmsg)
		log.Warning(msg)
		resp := response.StatusMsg(msg, start)
		return resp, nil
	}

	log.Debug("GetReportsByAddress:\n%+v\n", st)

	rpts, _ := data.D.FindAddress(st.Address, st.radius, st.maxResults)
	log.Debug(">>> rpts:\n%s\n", spew.Sdump(rpts))

	resp, _ := response.NewResponseReports(true, st.Start(), rpts)

	return resp, nil
}

type GetReportsByAddressType struct {
	Request_Type
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

func (st *GetReportsByAddressType) validate() string {
	var v validate

	st.radius = v.float("Radius", st.Radius)
	st.maxResults = v.int("MaxResults", st.MaxResults)
	st.includeDetails = v.bool("IncludeDetails", st.IncludeDetails)
	return v.errmsg
}

func (st GetReportsByAddressType) String() string {
	ls := new(logs.LogString)
	ls.AddS("GetReportsByAddress\n")
	ls.AddS(st.Request_Type.String())
	ls.AddF("Address \"%v\"\n", st.Address)
	ls.AddF("Radius %s/%v   MaxResults: %s/%v\n", st.Radius, st.radius, st.MaxResults, st.maxResults)
	ls.AddF("IncludeDetails: %v/%t\n", st.IncludeDetails, st.includeDetails)
	ls.AddF("Date Range: %v  to: %v \n", st.DateRangeStart, st.DateRangeEnd)
	return ls.Box(90)
}

// ==============================================================================================================================
//                                      GetReportsByLatLng
// ==============================================================================================================================

func GetReportsByLatLng(input string, start time.Time) (string, error) {
	st := new(GetReportsByLatLngType)
	err := xml.Unmarshal([]byte(input), st)
	if err != nil {
		msg := fmt.Sprintf("Unable to unmarshal GetReportsByLatLng request: %s", err)
		log.Warning(msg)
		return "", errors.New(msg)
	}

	st.start = start
	if errmsg := st.validate(); errmsg != "" {
		msg := fmt.Sprintf("Invalid request - %s", errmsg)
		log.Warning(msg)
		resp := response.StatusMsg(msg, start)
		return resp, nil
	}

	log.Debug("GetReportsByLatLng:\n%+v\n", st)

	rpts, _ := data.D.FindLL(st.latitude, st.longitude, st.radius, st.maxResults)
	log.Debug(">>> rpts:\n%s\n", spew.Sdump(rpts))

	resp, _ := response.NewResponseReports(true, st.Start(), rpts)

	return resp, nil
}

type GetReportsByLatLngType struct {
	Request_Type
	Latitude       string  `xml:"Latitude" json:"Latitude"`
	latitude       float64 //
	Longitude      string  `xml:"Longitude" json:"Longitude"`
	longitude      float64 //
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

func (st *GetReportsByLatLngType) validate() string {
	var v validate

	st.latitude = v.float("Latitude", st.Latitude)
	st.longitude = v.float("Longitude", st.Longitude)
	st.radius = v.float("Radius", st.Radius)
	st.maxResults = v.int("MaxResults", st.MaxResults)
	st.includeDetails = v.bool("IncludeDetails", st.IncludeDetails)
	return v.errmsg
}

func (st GetReportsByLatLngType) String() string {
	ls := new(logs.LogString)
	ls.AddS("GetReportsByLatLng\n")
	ls.AddS(st.Request_Type.String())
	ls.AddF("Loc \"%v:%v\"\n", st.latitude, st.longitude)
	ls.AddF("Radius %s/%v   MaxResults: %s/%v\n", st.Radius, st.radius, st.MaxResults, st.maxResults)
	ls.AddF("IncludeDetails: %v/%t\n", st.IncludeDetails, st.includeDetails)
	ls.AddF("Date Range: %v  to: %v \n", st.DateRangeStart, st.DateRangeEnd)
	return ls.Box(90)
}
