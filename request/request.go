package request

import (
	"CitySourcedAPI/data"
	"CitySourcedAPI/logs"

	"encoding/xml"
	"errors"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

var debug = true
var verbose = true
var log = logs.Log

// ==============================================================================================================================
//                                       PROCESS REQUEST
// ==============================================================================================================================

func ProcessRequest(input string) (string, error) {
	rt := new(Request_Type)
	if err := xml.Unmarshal([]byte(input), rt); err != nil {
		return "", err
	}
	if debug {
		log.Debug("Request received: %s", rt.ApiRequestType)
		log.Debug("Request:\n%s\n", spew.Sdump(*rt))
	}

	switch rt.ApiRequestType {
	case "CreateThreeOneOne":
		log.Debug("Processing CreateThreeOneOne request...")
		data, err := ProcessCreateThreeOneOne(input)
		if err != nil {
			log.Warning("ProcessCreateThreeOneOne failed - error: %s", err)
		}
		if debug {
			fmt.Println(spew.Sdump(data))
		}

	default:
		msg := fmt.Sprintf("Unknown request received: %s", rt.ApiRequestType)
		log.Warning(msg)
		return "", errors.New(msg)
	}

	return "", nil

}

// ==============================================================================================================================
//                                       REQUEST
// ==============================================================================================================================
type Request_Type struct {
	XMLName           xml.Name `xml:"CsRequest" json:"CsRequest"`
	ApiAuthKey        string   `xml:"ApiAuthKey" json:"ApiAuthKey"`
	ApiRequestType    string   `xml:"ApiRequestType" json:"ApiRequestType"`
	ApiRequestVersion string   `xml:"ApiRequestVersion" json:"ApiRequestVersion"`
}

// Check auth code.
func (r *Request_Type) auth() (ok bool) {
	ok = data.Auth(r.ApiAuthKey)
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
