package request

import (
	"CitySourcedAPI/logs"

	"encoding/xml"
	"errors"
	"fmt"
)

type KeyValuePair_Type struct {
	Value string `xml:",chardata"`
	Key   string `xml:"Key,attr"`
}

type CreateThreeOneOne_Type struct {
	Request_Type
	DateCreated       data.CustomTime     `xml:"DateCreated" json:"DateCreated"`
	DeviceType        string              `xml:"DeviceType" json:"DeviceType"`
	DeviceModel       string              `xml:"DeviceModel" json:"DeviceModel"`
	DeviceId          string              `xml:"DeviceId" json:"DeviceId"`
	RequestType       string              `xml:"RequestType" json:"RequestType"`
	RequestTypeId     string              `xml:"RequestTypeId" json:"RequestTypeId"`
	Latitude          float64             `xml:"Latitude" json:"Latitude"`
	Longitude         float64             `xml:"Longitude" json:"Longitude"`
	Directionality    string              `xml:"Directionality" json:"Directionality"`
	Description       string              `xml:"Description" json:"Description"`
	AuthorNameFirst   string              `xml:"AuthorNameFirst" json:"AuthorNameFirst"`
	AuthorNameLast    string              `xml:"AuthorNameLast" json:"AuthorNameLast"`
	AuthorEmail       string              `xml:"AuthorEmail" json:"AuthorEmail"`
	AuthorTelephone   string              `xml:"AuthorTelephone" json:"AuthorTelephone"`
	AuthorIsAnonymous bool                `xml:"AuthorIsAnonymous" json:"AuthorIsAnonymous"`
	KeyValuePairs     []KeyValuePair_Type `xml:"KeyValuePairs>KeyValuePair"`
}

func CreateThreeOneOne(input string) (st *CreateThreeOneOne_Type, err error) {
	st = new(CreateThreeOneOne_Type)
	err = xml.Unmarshal([]byte(input), st)
	if err != nil {
		msg := fmt.Sprintf("Unable to unmarshal CreateThreeOneOne request: %s", err)
		log.Warning(msg)
		return st, errors.New(msg)
	}
	return st, nil
}

func (s CreateThreeOneOne_Type) String() string {
	ls := new(logs.LogString)
	ls.AddS("CreateThreeOneOne_Type\n")
	ls.AddS(s.Request_Type.String())
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
