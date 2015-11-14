package data

import (
	"CitySourcedAPI/logs"
	"_sketches/spew"
	"encoding/xml"
	"errors"
	"fmt"
	"sort"
	"strconv"
)

const (
	dfltLatitude          float64 = 0.0
	dfltLongitude         float64 = 0.0
	dfltAuthorIsAnonymous         = true
	dfltVotes             int64   = 0
)

// ==============================================================================================================================
//                                      REPORT LIST
// ==============================================================================================================================

// ------------------------------- ReportList -------------------------------
type ReportList []*Report

func newReportList() ReportList {
	var x []*Report
	return x
}

func (rl *ReportList) Add(r *Report) error {
	log.Debug("[ReportList.Add] r:\n%s\n", spew.Sdump(r))
	*rl = append(*rl, r)
	return nil
}

func (rl *ReportList) AddBR(id int64, st *BaseReport) (*Report, error) {
	r := Report{
		ID:         id,
		BaseReport: *st,
	}
	log.Debug("[Add] st: type: %T\n%s\n", st, spew.Sdump(r))
	rl.Add(&r)
	return &r, nil
}

// ------------------------------- ReportListD -------------------------------
// Has "Distance" capabilities
type ReportListD struct {
	ReportList
	dist []float64
}

func NewReportListD() ReportListD {
	var x ReportListD
	x.ReportList = make([]*Report, 0)
	x.dist = make([]float64, 0)
	return x
}

func (rl *ReportListD) Add(r *Report, d float64) error {
	rl.ReportList.Add(r)
	rl.dist = append(rl.dist, d)
	log.Debug("[ReportListD.Add] r:\n%s\n", spew.Sdump(rl))
	return nil
}

func (rl *ReportListD) Len() int {
	return len(rl.dist)
}

func (rl *ReportListD) Less(i, j int) bool {
	return rl.dist[i] < rl.dist[j]
}

func (rl *ReportListD) Swap(i, j int) {
	rl.dist[i], rl.dist[j] = rl.dist[j], rl.dist[i]
	rl.ReportList[i], rl.ReportList[j] = rl.ReportList[j], rl.ReportList[i]
}

func (rl *ReportListD) Sort() {
	if len(rl.dist) > 0 {
		sort.Sort(rl)
	}
}

func (rl *ReportListD) Limit(n int64) {
	if (n > 0) && (n < int64(len(rl.ReportList))) {
		rl.ReportList = rl.ReportList[:n]
		rl.dist = rl.dist[:n]
	}
}

// ==============================================================================================================================
//                                      REPORT
// ==============================================================================================================================

// ------------------------------- Report_Type -------------------------------
type Report struct {
	XMLName xml.Name `xml:"Report" json:"Report"`
	ID      int64    `json:"Id" xml:"Id"`
	BaseReport
}

// Displays the contents of the Spec_Type custom type.
func (s Report) String() string {
	ls := new(logs.LogString)
	ls.AddS("Report\n")
	ls.AddF("ID: %v\n", s.ID)
	ls.AddS(s.BaseReport.String())
	return ls.Box(90)
}

// ==============================================================================================================================
//                                      BASE REPORT
// ==============================================================================================================================

// ------------------------------- BaseReport_Type -------------------------------
type BaseReport struct {
	DateCreated        CustomTime `json:"DateCreated" xml:"DateCreated"`
	DateUpdated        CustomTime `json:"DateUpdated" xml:"DateUpdated"`
	DeviceType         string     `json:"DeviceType" xml:"DeviceType"`
	DeviceModel        string     `json:"DeviceModel" xml:"DeviceModel"`
	DeviceID           string     `json:"DeviceId" xml:"DeviceId"`
	RequestType        string     `json:"RequestType" xml:"RequestType"`
	RequestTypeID      string     `json:"RequestTypeId" xml:"RequestTypeId"`
	ImageUrl           string     `json:"ImageUrl" xml:"ImageUrl"`
	ImageUrlXl         string     `json:"ImageUrlXl" xml:"ImageUrlXl"`
	ImageUrlLg         string     `json:"ImageUrlLg" xml:"ImageUrlLg"`
	ImageUrlMd         string     `json:"ImageUrlMd" xml:"ImageUrlMd"`
	ImageUrlSm         string     `json:"ImageUrlSm" xml:"ImageUrlSm"`
	ImageUrlXs         string     `json:"ImageUrlXs" xml:"ImageUrlXs"`
	City               string     `json:"City" xml:"City"`
	State              string     `json:"State" xml:"State"`
	ZipCode            string     `json:"ZipCode" xml:"ZipCode"`
	Latitude           string     `xml:"Latitude" json:"Latitude"`
	LatitudeV          float64    //
	Longitude          string     `xml:"Longitude" json:"Longitude"`
	LongitudeV         float64    //
	Directionality     string     `json:"Directionality" xml:"Directionality"`
	Description        string     `json:"Description" xml:"Description"`
	AuthorNameFirst    string     `json:"AuthorNameFirst" xml:"AuthorNameFirst"`
	AuthorNameLast     string     `json:"AuthorNameLast" xml:"AuthorNameLast"`
	AuthorEmail        string     `json:"AuthorEmail" xml:"AuthorEmail"`
	AuthorTelephone    string     `json:"AuthorTelephone" xml:"AuthorTelephone"`
	AuthorIsAnonymous  string     `xml:"AuthorIsAnonymous" json:"AuthorIsAnonymous"`
	AuthorIsAnonymousV bool       //
	UrlDetail          string     `json:"UrlDetail" xml:"UrlDetail"`
	UrlShortened       string     `json:"UrlShortened" xml:"UrlShortened"`
	Votes              string     `json:"Votes" xml:"Votes"`
	VotesV             int64      //
	StatusType         string     `json:"StatusType" xml:"StatusType"`
}

func (st *BaseReport) AuthIsAnon() bool {
	return st.AuthorIsAnonymousV
}

func (st *BaseReport) Validate() error {
	errmsg := ""

	// Longitude, Latitude - if there's a value, then convert it... otherwise
	// leave the st.LongitudeV and st.LatitudeV initialized to zero.
	if st.Latitude == "" {
		st.LatitudeV = dfltLatitude
	} else {
		x, err := strconv.ParseFloat(st.Latitude, 64)
		if err != nil {
			errmsg = errmsg + fmt.Sprintf("Invalid Latitude: %s\n", st.Latitude)
			st.LatitudeV = dfltLatitude
		}
		st.LatitudeV = x
	}

	if st.Longitude == "" {
		st.LongitudeV = dfltLongitude
	} else {
		x, err := strconv.ParseFloat(st.Longitude, 64)
		if err != nil {
			errmsg = errmsg + fmt.Sprintf("Invalid Longitude: %s\n", st.Longitude)
			st.LongitudeV = dfltLongitude
		}
		st.LongitudeV = x
	}

	// AuthorIsAnonymous
	if st.AuthorIsAnonymous == "" {
		st.AuthorIsAnonymousV = dfltAuthorIsAnonymous
	} else {
		x, err := strconv.ParseBool(st.AuthorIsAnonymous)
		if err != nil {
			errmsg = errmsg + fmt.Sprintf("Invalid AuthorIsAnonymous: %s\n", st.AuthorIsAnonymous)
			st.AuthorIsAnonymousV = dfltAuthorIsAnonymous
		}
		st.AuthorIsAnonymousV = x
	}

	// Votes
	if st.Votes == "" {
		st.VotesV = dfltVotes
	} else {
		x, err := strconv.ParseInt(st.Votes, 10, 64)
		if err != nil {
			errmsg = errmsg + fmt.Sprintf("Invalid Votes: %s\n", st.Votes)
			st.VotesV = dfltVotes
		}
		st.VotesV = x
	}

	if errmsg != "" {
		return errors.New(errmsg)
	}
	return nil
}

func (r *BaseReport) CalcDistance(rlat, rlon float64) float64 {
	return Distance(rlat, rlon, r.LatitudeV, r.LongitudeV)
}

// Displays the contents of the Spec_Type custom type.
func (s BaseReport) String() string {
	ls := new(logs.LogString)
	ls.AddS("Base Report\n")
	ls.AddF("DateCreated \"%v\"\n", s.DateCreated)
	ls.AddF("Device - type %s  model: %s  ID: %s\n", s.DeviceType, s.DeviceModel, s.DeviceID)
	ls.AddF("Request - type: %q  id: %q\n", s.RequestType, s.RequestTypeID)
	ls.AddF("Location - lat: %v  lon: %v  directionality: %q\n", s.LatitudeV, s.LongitudeV, s.Directionality)
	ls.AddF("          %s, %s   %s\n", s.City, s.State, s.ZipCode)
	ls.AddF("Votes: %d\n", s.VotesV)
	ls.AddF("Description: %q\n", s.Description)
	ls.AddF("Images - std: %s\n", s.ImageUrl)
	ls.AddF("          XS: %s\n", s.ImageUrlXs)
	ls.AddF("          SM: %s\n", s.ImageUrlSm)
	ls.AddF("          MD: %s\n", s.ImageUrlMd)
	ls.AddF("          LG: %s\n", s.ImageUrlLg)
	ls.AddF("          XL: %s\n", s.ImageUrlXl)
	ls.AddF("Author(anon: %t) %s %s  Email: %s  Tel: %s\n", s.AuthorIsAnonymousV, s.AuthorNameFirst, s.AuthorNameLast, s.AuthorEmail, s.AuthorTelephone)
	return ls.Box(80)
}
