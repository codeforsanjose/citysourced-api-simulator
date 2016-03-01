package response

import "CitySourcedAPI/data"

// ==============================================================================================================================
//                                      Reports
// ==============================================================================================================================

func prepResponse(src []*data.Report) []*Report {
	cr := make([]*Report, 0)
	for _, r := range src {
		cr = append(cr, convertReport(r))
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
	Description    string          `json:"Description" xml:"Description"`
	Votes          int64           `json:"Votes" xml:"Votes"`
	City           string          `json:"City" xml:"City"`
	State          string          `json:"State" xml:"State"`
	ZipCode        string          `json:"ZipCode" xml:"ZipCode"`
	Latitude       float64         `xml:"Latitude" json:"Latitude"`
	Longitude      float64         `xml:"Longitude" json:"Longitude"`
	Directionality string          `json:"Directionality" xml:"Directionality"`
	ImageUrl       string          `json:"ImageUrl" xml:"ImageUrl"`
	ImageUrlXl     string          `json:"ImageUrlXl" xml:"ImageUrlXl"`
	ImageUrlLg     string          `json:"ImageUrlLg" xml:"ImageUrlLg"`
	ImageUrlMd     string          `json:"ImageUrlMd" xml:"ImageUrlMd"`
	ImageUrlSm     string          `json:"ImageUrlSm" xml:"ImageUrlSm"`
	ImageUrlXs     string          `json:"ImageUrlXs" xml:"ImageUrlXs"`
	UrlDetail      string          `json:"UrlDetail" xml:"UrlDetail"`
	UrlShortened   string          `json:"UrlShortened" xml:"UrlShortened"`
	StatusType     string          `json:"StatusType" xml:"StatusType"`
	TicketSLA      string          `json:"TicketSla" xml:"TicketSla"`
	Comments       Comments        `json:"Comments,omitempty" xml:"Comments,omitempty"`
}

func convertReport(src *data.Report) *Report {
	r := Report{
		ID:           src.ID,
		DateCreated:  src.DateCreated,
		DateUpdated:  src.DateUpdated,
		RequestType:  src.RequestType,
		Description:  src.Description,
		Votes:        src.VotesV,
		City:         src.City,
		State:        src.State,
		ZipCode:      src.ZipCode,
		Latitude:     src.LatitudeV,
		Longitude:    src.LongitudeV,
		ImageUrl:     src.ImageUrl,
		ImageUrlXl:   src.ImageUrlXl,
		ImageUrlLg:   src.ImageUrlLg,
		ImageUrlMd:   src.ImageUrlMd,
		ImageUrlSm:   src.ImageUrlSm,
		ImageUrlXs:   src.ImageUrlXs,
		UrlDetail:    src.UrlDetail,
		UrlShortened: src.UrlShortened,
		StatusType:   src.StatusType,
		TicketSLA:    src.TicketSLA,
	}
	for _, c := range src.Comments {
		r.Comments.Comments = append(r.Comments.Comments, Comment{DateCreated: c.DateCreated, Comment: c.Comment})
	}
	r.Comments.CommentCount = len(src.Comments)
	return &r
}

type Comments struct {
	CommentCount int       `xml:"CommentCount"`
	Comments     []Comment `json:"Comment,omitempty" xml:"Comment,omitempty"`
}

type Comment struct {
	DateCreated data.CustomTime `json:"DateCreated" xml:"DateCreated"`
	Comment     string          `json:"Text" xml:"Text"`
}
