package request

import (
	// "CitySourcedAPI/common"
	"CitySourcedAPI/data"
	// "CitySourcedAPI/logs"

	"encoding/json"
	"encoding/xml"
	"fmt"
)

// ==============================================================================================================================
//                                       RESPONSE
// ==============================================================================================================================

type Response_Type struct {
	XMLName      xml.Name `xml:"CsResponse"`
	Message      string   `xml:"Message"`
	ResponseTime string   `xml:"ResponseTime"`
}

func (r *Response_Type) xml() (string, error) {
	b, err := xml.Marshal(r)
	return string(b), err
}

func (r *Response_Type) json() (string, error) {
	b, err := json.Marshal(r)
	return string(b), err
}

func NewResponseShort(message string, rtime float64) (*Response_Type, error) {
	rt := Response_Type{
		Message:      message,
		ResponseTime: fmt.Sprintf("%v Seconds", rtime),
	}
	return &rt, nil
}

// ==============================================================================================================================
//                                       REPORTS
// ==============================================================================================================================
type ResponseReports_Type struct {
	Response_Type
	Reports []Report_Type `xml:"Reports>Report"`
}

type Report_Type struct {
	XMLName xml.Name `xml:"Request" json:"Request"`
	data.Report_Type
}

// ==============================================================================================================================
//                                       Address
// ==============================================================================================================================
type ResponseAddress_Type struct {
	Response_Type
	Address string `xml:"Address"`
}
