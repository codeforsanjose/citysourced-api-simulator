package response

import (
	"CitySourcedAPI/data"
	"CitySourcedAPI/logs"

	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

const (
	XmlHeader string = "<?xml version=\"1.0\" encoding=\"utf-8\" ?>"
)

var (
	log = logs.Log
	responseMsg map[bool]string
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
	b, err := xml.MarshalIndent(r, "", "   ")
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
	Reports []*data.Report_Type `xml:"Reports>Report"`
}

func (r *ResponseReports_Type) xml() (string, error) {
	b, err := xml.MarshalIndent(r, "", "   ")
	return string(b), err
}

func (r *ResponseReports_Type) json() (string, error) {
	b, err := json.Marshal(r)
	return string(b), err
}

func NewResponseReports(success bool, rtime float64, reports []*data.Report_Type) (string, error) {
	rt := ResponseReports_Type{}
	rt.Message = responseMsg[success]
	rt.ResponseTime = fmt.Sprintf("%v Seconds", rtime)
	rt.Reports = reports
	log.Debug("rt: %s", spew.Sdump(rt))
	xmlout, err := rt.xml()
	xmlout = XmlHeader + xmlout
	return xmlout, err
}

// ==============================================================================================================================
//                                       Address
// ==============================================================================================================================
type ResponseAddress_Type struct {
	Response_Type
	Address string `xml:"Address"`
}

func init() {
	responseMsg = make(map[bool]string)
	responseMsg[true] = "Congratulations! The reports you requested are below."
	responseMsg[false] = "FAIL"
}