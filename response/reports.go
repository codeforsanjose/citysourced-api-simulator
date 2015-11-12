package response

import "CitySourcedAPI/data"

// ==============================================================================================================================
//                                      Reports
// ==============================================================================================================================

func ConvertReports(src []*data.Report) []*Report {
	cr := make([]*Report, 0)
	for _, r := range src {
		cr = append(cr, ConvertReport(r))
	}
	return cr
}

// ==============================================================================================================================
//                                      Report
// ==============================================================================================================================

type Report struct {
	ID             int64           `json:"Id" xml:"Id"`
	DateCreated    data.CustomTime `json:"DateCreated" xml:"DateCreated"`
	DateUpdated    data.CustomTime `json:"DateUpdated" xml:"DateUpdated"`
	RequestType    string          `json:"RequestType" xml:"RequestType"`
	ImageUrl       string          `json:"ImageUrl" xml:"ImageUrl"`
	ImageUrlXl     string          `json:"ImageUrlXl" xml:"ImageUrlXl"`
	ImageUrlLg     string          `json:"ImageUrlLg" xml:"ImageUrlLg"`
	ImageUrlMd     string          `json:"ImageUrlMd" xml:"ImageUrlMd"`
	ImageUrlSm     string          `json:"ImageUrlSm" xml:"ImageUrlSm"`
	ImageUrlXs     string          `json:"ImageUrlXs" xml:"ImageUrlXs"`
	City           string          `json:"City" xml:"City"`
	State          string          `json:"State" xml:"State"`
	ZipCode        string          `json:"ZipCode" xml:"ZipCode"`
	Latitude       string          `xml:"Latitude" json:"Latitude"`
	Longitude      string          `xml:"Longitude" json:"Longitude"`
	Directionality string          `json:"Directionality" xml:"Directionality"`
	UrlDetail      string          `json:"UrlDetail" xml:"UrlDetail"`
	UrlShortened   string          `json:"UrlShortened" xml:"UrlShortened"`
	StatusType     string          `json:"StatusType" xml:"StatusType"`
}

func ConvertReport(src *data.Report) *Report {
	r := Report{
		DateCreated:  src.DateCreated,
		DateUpdated:  src.DateUpdated,
		RequestType:  src.RequestType,
		ImageUrl:     src.ImageUrl,
		ImageUrlXl:   src.ImageUrlXl,
		ImageUrlLg:   src.ImageUrlLg,
		ImageUrlMd:   src.ImageUrlMd,
		ImageUrlSm:   src.ImageUrlSm,
		ImageUrlXs:   src.ImageUrlXs,
		City:         src.City,
		State:        src.State,
		ZipCode:      src.ZipCode,
		Latitude:     src.Latitude,
		Longitude:    src.Longitude,
		UrlDetail:    src.UrlDetail,
		UrlShortened: src.UrlShortened,
		StatusType:   src.StatusType,
	}
	return &r
}
