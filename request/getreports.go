package request

import (
	"CitySourcedAPI/data"
	"CitySourcedAPI/logs"
	"CitySourcedAPI/response"

	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
)

const (
	dfltFindAddressSearchRadius float64 = 100.0 // Meters
	dfltMaxResults              int64   = 10
	dfltIncludeDetails                  = true
)

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

func GetReportsByAddress(input string, start time.Time) (string, error) {
	st := new(GetReportsByAddressType)
	err := xml.Unmarshal([]byte(input), st)
	if err != nil {
		msg := fmt.Sprintf("Unable to unmarshal GetReportsByAddress request: %s", err)
		log.Warning(msg)
		return "", errors.New(msg)
	}
	st.start = start

	// Validate
	errmsg := ""

	// Convert Radius
	if st.Radius == "" {
		st.radius = dfltFindAddressSearchRadius
	} else {
		st.radius, err = strconv.ParseFloat(st.Radius, 64)
		if err != nil {
			errmsg = errmsg + fmt.Sprintf("Invalid search radius: %s\n", st.Radius)
		}
	}

	// Convert MaxResults
	if st.MaxResults == "" {
		st.maxResults = dfltMaxResults
	} else {
		st.maxResults, err = strconv.ParseInt(st.MaxResults, 10, 64)
		if err != nil {
			errmsg = errmsg + fmt.Sprintf("Invalid Max Results: %s\n", st.MaxResults)
		}
	}

	// Convert IncludeDetails
	if st.IncludeDetails == "" {
		st.includeDetails = dfltIncludeDetails
	} else {
		st.includeDetails, err = strconv.ParseBool(st.IncludeDetails)
		if err != nil {
			errmsg = errmsg + fmt.Sprintf("Invalid Include Details: %s\n", st.IncludeDetails)
		}
	}
	log.Debug("GetReportsByAddress:\n%+v\n", st)

	rpts, _ := data.D.FindAddress(st.Address, st.radius, st.maxResults)
	log.Debug(">>> rpts:\n%s\n", spew.Sdump(rpts))

	resp, _ := response.NewResponseReports(true, st.Start(), rpts)

	return resp, nil
}

func (s GetReportsByAddressType) String() string {
	ls := new(logs.LogString)
	ls.AddS("GetReportsByAddress\n")
	ls.AddS(s.Request_Type.String())
	ls.AddF("Address \"%v\"\n", s.Address)
	ls.AddF("Radius %s/%v   MaxResults: %s/%v\n", s.Radius, s.radius, s.MaxResults, s.maxResults)
	ls.AddF("IncludeDetails: %v/%t\n", s.IncludeDetails, s.includeDetails)
	ls.AddF("Date Range: %v  to: %v \n", s.DateRangeStart, s.DateRangeEnd)
	return ls.Box(90)
}
