package request

import (
	"CitySourcedAPI/logs"
	"CitySourcedAPI/data"
	_ "fmt"
	"encoding/xml"
)

var debug = true
var verbose = true
var log = logs.Log


type Request_Type struct {
	XMLName           xml.Name            `xml:"CsRequest"`
	ApiAuthKey        string              `xml:"ApiAuthKey"`
	ApiRequestType    string              `xml:"ApiRequestType"`
	ApiRequestVersion string              `xml:"ApiRequestVersion"`
}

// Check auth code.
func (r *Request_Type) auth() error {
	if r.ApiAuthKey != data.System
}

// Displays the contents of the Spec_Type custom type.
func (s Request_Type) String() string {
	ls := new(logs.LogString)
	ls.AddS("Request_Type\n")
	ls.AddF("Request - type: %s  ver: %s\n", s.ApiRequestType, s.ApiRequestVersion)
	return ls.BoxC(60)
}
