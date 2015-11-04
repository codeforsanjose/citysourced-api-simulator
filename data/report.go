package data

import (
	"CitySourcedAPI/common"
	"CitySourcedAPI/logs"
)

// ==============================================================================================================================
//                                      REPORTS
// ==============================================================================================================================

// ------------------------------- Report_Type -------------------------------
type Report_Type struct {
	Id                int64             `json:"Id" xml:"Id"`
	DateCreated       common.CustomTime `json:"DateCreated" xml:"DateCreated"`
	DateUpdated       common.CustomTime `json:"DateUpdated" xml:"DateUpdated"`
	DeviceType        string            `json:"DeviceType" xml:"DeviceType"`
	DeviceModel       string            `json:"DeviceModel" xml:"DeviceModel"`
	DeviceId          string            `json:"DeviceId" xml:"DeviceId"`
	RequestType       string            `json:"RequestType" xml:"RequestType"`
	RequestTypeId     string            `json:"RequestTypeId" xml:"RequestTypeId"`
	ImageUrl          string            `json:"ImageUrl" xml:"ImageUrl"`
	ImageUrlXl        string            `json:"ImageUrlXl" xml:"ImageUrlXl"`
	ImageUrlLg        string            `json:"ImageUrlLg" xml:"ImageUrlLg"`
	ImageUrlMd        string            `json:"ImageUrlMd" xml:"ImageUrlMd"`
	ImageUrlSm        string            `json:"ImageUrlSm" xml:"ImageUrlSm"`
	ImageUrlXs        string            `json:"ImageUrlXs" xml:"ImageUrlXs"`
	City              string            `json:"City" xml:"City"`
	State             string            `json:"State" xml:"State"`
	ZipCode           string            `json:"ZipCode" xml:"ZipCode"`
	Latitude          float64           `json:"Latitude" xml:"Latitude"`
	Longitude         float64           `json:"Longitude" xml:"Longitude"`
	Directionality    string            `json:"Directionality" xml:"Directionality"`
	Description       string            `json:"Description" xml:"Description"`
	AuthorNameFirst   string            `json:"AuthorNameFirst" xml:"AuthorNameFirst"`
	AuthorNameLast    string            `json:"AuthorNameLast" xml:"AuthorNameLast"`
	AuthorEmail       string            `json:"AuthorEmail" xml:"AuthorEmail"`
	AuthorTelephone   string            `json:"AuthorTelephone" xml:"AuthorTelephone"`
	AuthorIsAnonymous bool              `json:"AuthorIsAnonymous" xml:"AuthorIsAnonymous"`
	UrlDetail         string            `json:"UrlDetail" xml:"UrlDetail"`
	UrlShortened      string            `json:"UrlShortened" xml:"UrlShortened"`
	StatusType        string            `json:"StatusType" xml:"StatusType"`
}

func (r *Report_Type) Distance(rlat, rlon float64) float64 {
	return Distance(rlat, rlon, r.Latitude, r.Longitude)
}

// Displays the contents of the Spec_Type custom type.
func (s Report_Type) String() string {
	ls := new(logs.LogString)
	ls.AddS("Report\n")
	ls.AddF("Id: %v\n", s.Id)
	ls.AddF("DateCreated \"%v\"\n", s.DateCreated)
	ls.AddF("Device - type %s  model: %s  Id: %s\n", s.DeviceType, s.DeviceModel, s.DeviceId)
	ls.AddF("Request - type: %q  id: %q\n", s.RequestType, s.RequestTypeId)
	ls.AddF("Location - lat: %v  lon: %v  directionality: %q\n", s.Latitude, s.Longitude, s.Directionality)
	ls.AddF("          %s, %s   %s\n", s.City, s.State, s.ZipCode)
	ls.AddF("Description: %q\n", s.Description)
	ls.AddF("Images - std: %s\n", s.ImageUrl)
	ls.AddF("          XS: %s\n", s.ImageUrlXs)
	ls.AddF("          SM: %s\n", s.ImageUrlSm)
	ls.AddF("          MD: %s\n", s.ImageUrlMd)
	ls.AddF("          LG: %s\n", s.ImageUrlLg)
	ls.AddF("          XL: %s\n", s.ImageUrlXl)
	ls.AddF("Author(anon: %t) %s %s  Email: %s  Tel: %s\n", s.AuthorIsAnonymous, s.AuthorNameFirst, s.AuthorNameLast, s.AuthorEmail, s.AuthorTelephone)
	return ls.Box(90)
}
