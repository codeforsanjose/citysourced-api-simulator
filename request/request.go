package request

import (
	"CitySourcedAPI/config"
	"CitySourcedAPI/logs"
	"CitySourcedAPI/response"
	"errors"

	"encoding/xml"
	// "errors"
	"fmt"
	"reflect"
	"time"
)

var (
	log          = logs.Log
	typeRegistry = make(map[string]reflect.Type)
)

// ==============================================================================================================================
//                                       INIT
// ==============================================================================================================================

func init() {

	typeRegistry["CreateThreeOneOne"] = reflect.TypeOf(CreateThreeOneOne{})
	typeRegistry["GetReportsByAddress"] = reflect.TypeOf(GetReportsByAddress{})
	typeRegistry["GetReportsByLatLng"] = reflect.TypeOf(GetReportsByLatLng{})
	typeRegistry["GetReportsByDeviceId"] = reflect.TypeOf(GetReportsByDeviceID{})
	typeRegistry["GetReportsByZipCode"] = reflect.TypeOf(GetReportsByZipCode{})
	typeRegistry["GetReport"] = reflect.TypeOf(GetReport{})
}

// ==============================================================================================================================
//                                       PROCESS REQUEST
// ==============================================================================================================================

type Processor interface {
	Validate(time.Time) string
	Run() (string, error)
}

func Process(input string, start time.Time) (string, error) {
	rt, e := newRequest(input, start)
	if e != nil {
		msg := fmt.Sprintf("Error while parsing the request: %q", e)
		log.Error("%s", msg)
		return response.StatusMsg(msg, start), errors.New(msg)
	}

	if ok := rt.auth(); !ok {
		msg := "Invalid auth code"
		log.Warning("%s", msg)
		return response.StatusMsg(msg, start), errors.New(msg)
	}

	// Create an instance of the request struct
	svcName, ok := typeRegistry[rt.ApiRequestType]
	if !ok {
		msg := fmt.Sprintf("Unknown request received: %s", rt.ApiRequestType)
		log.Warning(msg)
		return response.StatusMsg(msg, start), e
	}
	svc := reflect.New(svcName).Interface()

	// Unmarshal into it
	err := xml.Unmarshal([]byte(input), svc)
	if err != nil {
		msg := fmt.Sprintf("Unable to unmarshal GetReportsByLatLng request: %s", err)
		log.Warning(msg)
		return "", errors.New(msg)
	}

	// Validate
	errmsg := svc.(Processor).Validate(start)
	if errmsg != "" {
		msg := fmt.Sprintf("Invalid request - %s", errmsg)
		log.Warning(msg)
		resp := response.StatusMsg(msg, start)
		return resp, nil

	}

	// Run
	rsp, err := svc.(Processor).Run()

	log.Debug("Response:\n%s\n", rsp)
	if err != nil {
		log.Warning("Request failed - error: %s", err)
	}

	return rsp, err
}

// ==============================================================================================================================
//                                       REQUEST
// ==============================================================================================================================
func newRequest(input string, start time.Time) (*Request, error) {
	log.Debug("New Request: \n%s\n", input)
	rt := new(Request)
	if err := xml.Unmarshal([]byte(input), rt); err != nil {
		log.Warning("Request XML cannot be unmarshaled: %s", err)
		return nil, err
	}
	rt.start = start
	log.Debug("rt:\n%+v", rt)
	return rt, nil
}

type Request struct {
	start             time.Time
	XMLName           xml.Name `xml:"CsRequest" json:"CsRequest"`
	ApiAuthKey        string   `xml:"ApiAuthKey" json:"ApiAuthKey"`
	ApiRequestType    string   `xml:"ApiRequestType" json:"ApiRequestType"`
	ApiRequestVersion string   `xml:"ApiRequestVersion" json:"ApiRequestVersion"`
}

func (r *Request) Start() time.Time {
	return r.start
}

func (r *Request) SetStart(start time.Time) {
	r.start = start
}

func (r *Request) auth() (ok bool) {
	ok = config.Auth(r.ApiAuthKey)
	return ok
}

func (r Request) String() string {
	ls := new(logs.LogString)
	ls.AddS("Request_Type\n")
	ls.AddF("Start: %v\n", r.start)
	ls.AddF("Request - type: %s  ver: %s\n", r.ApiRequestType, r.ApiRequestVersion)
	return ls.BoxC(60)
}
