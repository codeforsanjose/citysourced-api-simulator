package response

import (
	"CitySourcedAPI/data"
	"CitySourcedAPI/logs"

	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"time"
)

const (
	XmlHeader string = "<?xml version=\"1.0\" encoding=\"utf-8\" ?>\n"
)

var (
	log         = logs.Log
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

func StatusMsg(message string, start time.Time) string {
	rt := Response_Type{
		Message:      message,
		ResponseTime: fmt.Sprintf("%v Seconds", time.Since(start).Seconds()),
	}
	xmlout, _ := rt.xml()
	xmlout = XmlHeader + xmlout
	return xmlout
}

// ==============================================================================================================================
//                                       REPORTS
// ==============================================================================================================================
type ResponseReport_Type struct {
	Response_Type
	Reports ResponseReports_Type `xml:"Reports"`
}

func (r *ResponseReport_Type) xml() (string, error) {
	b, err := xml.MarshalIndent(r, "", "   ")
	return string(b), err
}

func (r *ResponseReport_Type) json() (string, error) {
	b, err := json.Marshal(r)
	return string(b), err
}

type ResponseReports_Type struct {
	ReportCount int            `xml:"ReportCount"`
	Reports     []*data.Report `xml:"Report"`
}

func NewResponseReports(success bool, start time.Time, reports []*data.Report) (string, error) {
	rt := ResponseReport_Type{}
	rt.Message = responseMsg[success]
	rt.ResponseTime = fmt.Sprintf("%v Seconds", time.Since(start).Seconds())

	rts := ResponseReports_Type{}
	rts.ReportCount = len(reports)
	rts.Reports = reports
	rt.Reports = rts

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
