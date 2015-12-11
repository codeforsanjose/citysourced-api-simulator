package request

import (
	"CitySourcedAPI/data"
	"CitySourcedAPI/logs"
	"CitySourcedAPI/response"

	"time"
)

// ==============================================================================================================================
//                                      GetReportsByLatLng
// ==============================================================================================================================

type CreateThreeOneOne struct {
	Request
	Processor
	data.BaseReport
	KeyValuePairs []KeyValuePair `xml:"KeyValuePairs>KeyValuePair"`
}

type KeyValuePair struct {
	Value string `xml:",chardata"`
	Key   string `xml:"Key,attr"`
}

func (st *CreateThreeOneOne) validate() error {
	return st.BaseReport.Validate()
}

func (st *CreateThreeOneOne) Validate(start time.Time) string {
	var v validate
	st.start = start
	st.LatitudeV = v.float("Latitude", st.Latitude)
	st.LongitudeV = v.float("Longitude", st.Longitude)
	st.AuthorIsAnonymousV = v.bool("AuthorIsAnonymous", st.AuthorIsAnonymous)
	return v.errmsg
}

func (st *CreateThreeOneOne) Run() (string, error) {
	reportid, authorid, _ := data.Append(st.BaseReport)
	return response.NewResponseReport("New report created OK", st.start, reportid, authorid)
}

func (s CreateThreeOneOne) String() string {
	ls := new(logs.LogString)
	ls.AddS("CreateThreeOneOne_Type\n")
	ls.AddS(s.Request.String())
	ls.AddF("DateCreated \"%v\"\n", s.DateCreated)
	ls.AddF("Device - type %s  model: %s  Id: %s\n", s.DeviceType, s.DeviceModel, s.DeviceID)
	ls.AddF("Request - type: %q  id: %q\n", s.RequestType, s.RequestTypeID)
	ls.AddF("Location - lat: %v  lon: %v  directionality: %q\n", s.Latitude, s.Longitude, s.Directionality)
	ls.AddF("Description: %q\n", s.Description)
	ls.AddF("Author(anon: %t) %s %s  Email: %s  Tel: %s\n", s.AuthorIsAnonymous, s.AuthorNameFirst, s.AuthorNameLast, s.AuthorEmail, s.AuthorTelephone)
	for _, v := range s.KeyValuePairs {
		ls.AddF("   %s: %s\n", v.Key, v.Value)
	}
	return ls.Box(90)
}
