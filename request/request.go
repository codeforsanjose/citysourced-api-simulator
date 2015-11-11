package request

import (
	"CitySourcedAPI/config"
	"CitySourcedAPI/logs"
	"CitySourcedAPI/response"

	"encoding/xml"
	// "errors"
	"fmt"
	"time"
)

var (
	log        = logs.Log
	processors map[string]func(string, time.Time) (string, error)
)

// ==============================================================================================================================
//                                       PROCESS REQUEST
// ==============================================================================================================================

func Process(input string, start time.Time) (rsp string, err error) {
	rt, e := parse(input, start)
	if e != nil {
		return response.StatusMsg("Unable to parse request XML", start), e
	}

	if ok := rt.auth(); !ok {
		return response.StatusMsg("Invalid Auth code", start), e
	}

	f, ok := processors[rt.ApiRequestType]
	if !ok {
		msg := fmt.Sprintf("Unknown request received: %s", rt.ApiRequestType)
		log.Warning(msg)
		return response.StatusMsg(msg, start), e
	}

	rsp, err = f(input, start)

	log.Debug("Response:\n%s\n", rsp)
	if err != nil {
		log.Warning("Request failed - error: %s", err)
	}

	return rsp, nil

}

// ==============================================================================================================================
//                                       REQUEST
// ==============================================================================================================================
func parse(input string, start time.Time) (*Request_Type, error) {
	log.Debug("New Request: \n%s\n", input)
	rt := new(Request_Type)
	if err := xml.Unmarshal([]byte(input), rt); err != nil {
		log.Warning("Request XML cannot be unmarshaled: %s", err)
		return nil, err
	}
	rt.start = start
	log.Debug("rt:\n%+v", rt)
	return rt, nil
}

type Request_Type struct {
	start             time.Time
	XMLName           xml.Name `xml:"CsRequest" json:"CsRequest"`
	ApiAuthKey        string   `xml:"ApiAuthKey" json:"ApiAuthKey"`
	ApiRequestType    string   `xml:"ApiRequestType" json:"ApiRequestType"`
	ApiRequestVersion string   `xml:"ApiRequestVersion" json:"ApiRequestVersion"`
}

func (r *Request_Type) Start() time.Time {
	return r.start
}

func (r *Request_Type) auth() (ok bool) {
	ok = config.Auth(r.ApiAuthKey)
	if !ok {
		msg := "Invalid auth code."
		log.Warning(msg)
	}
	return ok
}

func (r Request_Type) String() string {
	ls := new(logs.LogString)
	ls.AddS("Request_Type\n")
	ls.AddF("Start: %v\n", r.start)
	ls.AddF("Request - type: %s  ver: %s\n", r.ApiRequestType, r.ApiRequestVersion)
	return ls.BoxC(60)
}

// ==============================================================================================================================
//                                       INIT
// ==============================================================================================================================

func init() {
	processors = make(map[string]func(string, time.Time) (string, error))

	processors["CreateThreeOneOne"] = CreateThreeOneOne
	processors["GetReportsByAddress"] = GetReportsByAddress
	processors["GetReportsByLatLng"] = GetReportsByLatLng
}
