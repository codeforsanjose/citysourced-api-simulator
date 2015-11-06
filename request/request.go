package request

import (
	"CitySourcedAPI/config"
	"CitySourcedAPI/logs"

	"encoding/xml"
	"errors"
	"fmt"
)

var (
	log = logs.Log
)

// ==============================================================================================================================
//                                       PROCESS REQUEST
// ==============================================================================================================================

func Process(input string) (response string, err error) {
	rt, e := newRequest(input)
	if e != nil {
		return "", e
	}

	switch rt.ApiRequestType {
	case "CreateThreeOneOne":
		response, err = CreateThreeOneOne(input)

	case "GetReportsByAddress":
		response, err = GetReportsByAddress(input)

	default:
		msg := fmt.Sprintf("Unknown request received: %s", rt.ApiRequestType)
		log.Warning(msg)
		return "", errors.New(msg)
	}

	log.Debug("Response:\n%s\n", response)
	if err != nil {
		log.Warning("CreateThreeOneOne failed - error: %s", err)
	}

	return "", nil

}

// ==============================================================================================================================
//                                       REQUEST
// ==============================================================================================================================
func newRequest(input string) (*Request_Type, error) {
	log.Debug("New Request: \n%s\n", input)
	rt := new(Request_Type)
	if err := xml.Unmarshal([]byte(input), rt); err != nil {
		log.Warning("Request XML cannot be unmarshaled: %s", err)
		return nil, err
	}
	log.Debug("rt:\n%+v", rt)
	return rt, nil
}

type Request_Type struct {
	XMLName           xml.Name `xml:"CsRequest" json:"CsRequest"`
	ApiAuthKey        string   `xml:"ApiAuthKey" json:"ApiAuthKey"`
	ApiRequestType    string   `xml:"ApiRequestType" json:"ApiRequestType"`
	ApiRequestVersion string   `xml:"ApiRequestVersion" json:"ApiRequestVersion"`
}

// Check auth code.
func (r *Request_Type) auth() (ok bool) {
	ok = config.Auth(r.ApiAuthKey)
	if !ok {
		msg := "Invalid auth code."
		log.Warning(msg)
	}
	return ok
}

// Displays the contents of the Spec_Type custom type.
func (s Request_Type) String() string {
	ls := new(logs.LogString)
	ls.AddS("Request_Type\n")
	ls.AddF("Request - type: %s  ver: %s\n", s.ApiRequestType, s.ApiRequestVersion)
	return ls.BoxC(60)
}
