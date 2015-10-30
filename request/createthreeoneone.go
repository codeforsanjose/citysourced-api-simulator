package request

import (
	"CitySourcedAPI/common"
	"CitySourcedAPI/logs"
	"encoding/xml"
	_ "fmt"
)

type KeyValuePair_Type struct {
	Value string `xml:",chardata"`
	Key   string `xml:"Key,attr"`
}

type CreateThreeOneOne_Type struct {
	XMLName           xml.Name            `xml:"CsRequest"`
	ApiAuthKey        string              `xml:"ApiAuthKey"`
	ApiRequestType    string              `xml:"ApiRequestType"`
	ApiRequestVersion string              `xml:"ApiRequestVersion"`
	DateCreated       common.CustomTime   `xml:"DateCreated"`
	DeviceType        string              `xml:"DeviceType"`
	DeviceModel       string              `xml:"DeviceModel"`
	DeviceId          string              `xml:"DeviceId"`
	RequestType       string              `xml:"RequestType"`
	RequestTypeId     string              `xml:"RequestTypeId"`
	Latitude          float64             `xml:"Latitude"`
	Longitude         float64             `xml:"Longitude"`
	Directionality    string              `xml:"Directionality"`
	Description       string              `xml:"Description"`
	AuthorNameFirst   string              `xml:"AuthorNameFirst"`
	AuthorNameLast    string              `xml:"AuthorNameLast"`
	AuthorEmail       string              `xml:"AuthorEmail"`
	AuthorTelephone   string              `xml:"AuthorTelephone"`
	AuthorIsAnonymous bool                `xml:"AuthorIsAnonymous"`
	KeyValuePairs     []KeyValuePair_Type `xml:"KeyValuePairs>KeyValuePair"`
}

func NewCreateThreeOneOne(input string) (st *CreateThreeOneOne_Type, err error) {
	st = new(CreateThreeOneOne_Type)
	err = xml.Unmarshal([]byte(input), st)
	return st, err
}

// Displays the contents of the Spec_Type custom type.
func (s CreateThreeOneOne_Type) String() string {
	ls := new(logs.LogString)
	ls.AddS("CreateThreeOneOne_Type\n")
	ls.AddF("AuthKey: %q\n", s.ApiAuthKey)
	ls.AddF("Request - type: %s  ver: %s\n", s.ApiRequestType, s.ApiRequestVersion)
	ls.AddF("DateCreated \"%v\"\n", s.DateCreated)
	ls.AddF("Device - type %s  model: %s  Id: %s\n", s.DeviceType, s.DeviceModel, s.DeviceId)
	ls.AddF("Request - type: %q  id: %q\n", s.RequestType, s.RequestTypeId)
	ls.AddF("Location - lat: %v  lon: %v  directionality: %q\n", s.Latitude, s.Longitude, s.Directionality)
	ls.AddF("Description: %q\n", s.Description)
	ls.AddF("Author(anon: %t) %s %s  Email: %s  Tel: %s\n", s.AuthorIsAnonymous, s.AuthorNameFirst, s.AuthorNameLast, s.AuthorEmail, s.AuthorTelephone)
	for _, v := range s.KeyValuePairs {
		ls.AddF("   %s: %s\n", v.Key, v.Value)
	}
	return ls.Box(90)
}
