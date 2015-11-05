package request

import (
	"CitySourcedAPI/logs"

	"encoding/xml"
	"errors"
	"fmt"
)

type GetReportsByAddress_Type struct {
	Request_Type
	Address        string          `xml:"Address" json:"Address"`
	Radius         float64         `xml:"Radius" json:"Radius"`
	MaxResults     int64           `xml:"MaxResults" json:"MaxResults"`
	IncludeDetails bool            `xml:"IncludeDetails" json:"IncludeDetails"`
	DateRangeStart data.CustomTime `xml:"DateRangeStart" json:"DateRangeStart"`
	DateRangeEnd   data.CustomTime `xml:"DateRangeEnd" json:"DateRangeEnd"`
	CurrentStatus  string          `xml:"CurrentStatus" json:"CurrentStatus"`
}

func GetReportsByAddress(input string) (st *GetReportsByAddress_Type, err error) {
	st = new(GetReportsByAddress_Type)
	err = xml.Unmarshal([]byte(input), st)
	if err != nil {
		msg := fmt.Sprintf("Unable to unmarshal GetReportsByAddress request: %s", err)
		log.Warning(msg)
		return st, errors.New(msg)
	}
	return st, nil
}

func (s GetReportsByAddress_Type) String() string {
	ls := new(logs.LogString)
	ls.AddS("CreateThreeOneOne_Type\n")
	ls.AddS(s.Request_Type.String())
	ls.AddF("Address \"%v\"\n", s.Address)
	ls.AddF("Radius %v   MaxResults: %v\n", s.Radius, s.MaxResults)
	ls.AddF("IncludeDetails: %t\n", s.IncludeDetails)
	ls.AddF("Date Range: %v  to: %v \n", s.DateRangeStart, s.DateRangeEnd)
	return ls.Box(90)
}
