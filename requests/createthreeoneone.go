package requests

import (
	"encoding/xml"
	_ "fmt"
)

type KeyValuePair_Type struct {
	Value string `xml:",chardata"`
	Key   string `xml:"Key,attr"`
}

type CreateMobileDevice_Type struct {
	XMLName           xml.Name            `xml:"CsRequest"`
	ApiAuthKey        string              `xml:"ApiAuthKey"`
	ApiRequestType    string              `xml:"ApiRequestType"`
	ApiRequestVersion string              `xml:"ApiRequestVersion"`
	DateCreated       string              `xml:"DateCreated"`
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

func NewCreateMobileDevice(input string) (st *CreateMobileDevice_Type, err error) {
	st = new(CreateMobileDevice_Type)
	err = xml.Unmarshal([]byte(input), st)
	return st, err
}

/*

// Displays the contents of the Spec_Type custom type.
func (s Spec_Type) String() string {
	ls := new(logs.LogString)
	ls.AddF("[%s]\n", s.name)
	ls.AddF("Indexes - timestamp: %d   uniqueId: %d\n", s.timestampIndex, s.uniqueIdIndex)
	for _, col := range s.colOrder {
		ls.AddF("   %s\n", s.columns[col])
	}
	return ls.Box(90)
}

*/