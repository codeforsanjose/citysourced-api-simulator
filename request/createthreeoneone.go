package request

import (
	"CitySourcedAPI/data"
	"CitySourcedAPI/logs"

	"encoding/xml"
	"errors"
	"fmt"
	"time"
)

const (
	dfltLatitude          float64 = 0.0
	dfltLongitude         float64 = 0.0
	dfltAuthorIsAnonymous         = true
)

type KeyValuePair_Type struct {
	Value string `xml:",chardata"`
	Key   string `xml:"Key,attr"`
}

type CreateThreeOneOne_Type struct {
	Request_Type
	data.BaseReport_Type
	KeyValuePairs []KeyValuePair_Type `xml:"KeyValuePairs>KeyValuePair"`
}

func (st *CreateThreeOneOne_Type) Validate() error {
	return st.BaseReport_Type.Validate()
}

func CreateThreeOneOne(input string, start time.Time) (string, error) {
	st := new(CreateThreeOneOne_Type)
	if err := xml.Unmarshal([]byte(input), st); err != nil {
		msg := fmt.Sprintf("Unable to unmarshal CreateThreeOneOne request: %s", err)
		log.Warning(msg)
		return "", errors.New(msg)
	}
	st.start = start

	if err := st.Validate(); err != nil {
		return "", err
	}

	log.Debug("CreateThreeOneOne: \n%+v\n", st)

	data.D.AddReport(st.BaseReport_Type)

	return "", nil
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
