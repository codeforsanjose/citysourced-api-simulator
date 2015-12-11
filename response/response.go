package response

import (
	"CitySourcedAPI/data"
	"CitySourcedAPI/logs"

	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

const (
	XmlHeader string = "<?xml version=\"1.0\" encoding=\"utf-8\" ?>\n"
)

var (
	log         = logs.Log
	responseMsg map[bool]string
)

func StatusMsg(message string, start time.Time) string {
	rt := Response{
		Message:      message,
		ResponseTime: fmt.Sprintf("%v Seconds", time.Since(start).Seconds()),
	}
	xmlout, _ := rt.xml()
	xmlout = XmlHeader + xmlout
	return xmlout
}

// ==============================================================================================================================
//                                       RESPONSE
// ==============================================================================================================================

type Response struct {
	XMLName      xml.Name `xml:"CsResponse"`
	Message      string   `xml:"Message"`
	ResponseTime string   `xml:"ResponseTime"`
}

func (r *Response) xml() (string, error) {
	b, err := xml.MarshalIndent(r, "", "   ")
	return string(b), err
}

func (r *Response) json() (string, error) {
	b, err := json.Marshal(r)
	return string(b), err
}

// ==============================================================================================================================
//                                       NEW REPORT
// ==============================================================================================================================

type ResponseNewReport struct {
	XMLName      xml.Name `xml:"CsResponse"`
	Message      string   `xml:"Message"`
	ReportID     int64    `xml:"ReportId"`
	AuthorID     int64    `xml:"AuthorId"`
	ResponseTime string   `xml:"ResponseTime"`
}

func (r *ResponseNewReport) xml() (string, error) {
	b, err := xml.MarshalIndent(r, "", "   ")
	return string(b), err
}

func (r *ResponseNewReport) json() (string, error) {
	b, err := json.Marshal(r)
	return string(b), err
}

func NewResponseReport(message string, start time.Time, reportID, authorID int64) (string, error) {
	rt := ResponseNewReport{
		Message:      message,
		ReportID:     reportID,
		AuthorID:     authorID,
		ResponseTime: fmt.Sprintf("%v Seconds", time.Since(start).Seconds()),
	}
	xmlout, _ := rt.xml()
	xmlout = XmlHeader + xmlout
	return xmlout, nil
}

// ==============================================================================================================================
//                                       REPORTS
// ==============================================================================================================================
type ResponseReport struct {
	Response
	Reports ResponseReports `xml:"Reports"`
}

func (r *ResponseReport) xml() (string, error) {
	b, err := xml.MarshalIndent(r, "", "   ")
	return string(b), err
}

func (r *ResponseReport) json() (string, error) {
	b, err := json.Marshal(r)
	return string(b), err
}

type ResponseReports struct {
	ReportCount int       `xml:"ReportCount"`
	Reports     []*Report `xml:"Report"`
}

func NewResponseReports(success bool, start time.Time, reports []*data.Report) (string, error) {
	rt := ResponseReport{}
	rt.Message = responseMsg[success]
	rt.ResponseTime = fmt.Sprintf("%v Seconds", time.Since(start).Seconds())

	rts := ResponseReports{}
	rts.ReportCount = len(reports)
	rts.Reports = prepResponse(reports)
	rt.Reports = rts

	log.Debug("rt: %s", spew.Sdump(rt))
	xmlout, err := rt.xml()
	xmlout = XmlHeader + xmlout
	return xmlout, err
}

// ==============================================================================================================================
//                                       Address
// ==============================================================================================================================
type ResponseAddress struct {
	Response
	Address string `xml:"Address"`
}

func init() {
	responseMsg = make(map[bool]string)
	responseMsg[true] = "Congratulations! The reports you requested are below."
	responseMsg[false] = "FAIL"
}
