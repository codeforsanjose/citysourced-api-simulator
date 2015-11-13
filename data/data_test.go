package data_test

import (
	"CitySourcedAPI/data"
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestReadDataInvalidPath(t *testing.T) {
	fmt.Println("\n>>>>>>>>>>>>>>>>>>> TestReadDataInvalidPath <<<<<<<<<<<<<<<<<<<<<<<<<<")
	// Load Data
	if err := data.Init("../dataxxx.json"); err == nil {
		t.Errorf("Attempting to load an invalid file should have caused an error: %q", err)
	}
	fmt.Println("   (Should have just received a CRIT error 'Failed to open...')")
}

func TestReadDataInvalidJSON(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestReadDataInvalidJSON <<<<<<<<<<<<<<<<<<<<<<<<<<")
	// Load Data
	if err := data.Init("tests/data_faulty.json"); err == nil {
		t.Errorf("Attempting to load a faulty file should have caused an error: %q", err)
	}
	fmt.Println("    (Should have just received a CRIT error 'Invalid JSON...')")
}

func TestReadData(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestReadData <<<<<<<<<<<<<<<<<<<<<<<<<<")
	// Load Data
	if err := data.Init("../data.json"); err != nil {
		t.Errorf("Error %q occurred when loading the data.", err)
	}
	fmt.Printf("%v", data.D.Display())

}

// NOTE: if the test input data is changed, this must be updated!
func TestReportValidity(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestReportValidity <<<<<<<<<<<<<<<<<<<<<<<<<<")
	// Check lastID
	if lID := data.D.LastID(); lID != 102 {
		t.Errorf("LastId: %v is incorrect.", lID)
	}

	// Make sure we've got the data we think we should have - check random data:
	r, e := data.D.FindID(101)
	if e != nil {
		t.Errorf("FindId failed: %q.", e)
	}
	fmt.Printf("Id 100 - type: %T.\n", r)

	var cdval data.CustomTime
	cdval.UnmarshalText([]byte("2015-10-20T13:45:30"))
	if r.DateCreated != cdval {
		t.Errorf("Invalid DateCreated: %s  should be: %s", r.DateCreated, cdval)
	}
	sval := "2222"
	if r.DeviceID != sval {
		t.Errorf("Invalid DeviceId: %s  should be: %s", r.DeviceID, sval)
	}
	fval := -121.886329
	if r.LongitudeV != fval {
		t.Errorf("Invalid Longitude: %v  should be: %v", r.LongitudeV, fval)
	}

	bval := true
	if r.AuthIsAnon() != bval {
		t.Errorf("Invalid AuthorIsAnonymous: %v  should be: %v", r.AuthIsAnon(), bval)
	}
}

// NOTE: if the test input data is changed, this must be updated!
func TestCommentValidity(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestCommentValidity <<<<<<<<<<<<<<<<<<<<<<<<<<")
	// Check lastID
	if lID := data.LastCommentID(); lID != 3 {
		t.Errorf("LastId: %v is incorrect.", lID)
	}

	// Make sure we've got the data we think we should have - check random data:
	r, e := data.FindReportComments(101)
	fmt.Println(r)
	if e != nil {
		t.Errorf("FindId failed: %q.", e)
	}
	if l := len(r); l != 1 {
		t.Errorf("Size of comments for report 101 is %d should be %d.", l, 2)
	}

	sExp := "The second test comment for report 101"
	if cmt := r[1].Comment; cmt != sExp {
		t.Errorf("The text of the second comment for report 101 is %q should be %q.", cmt, sExp)
	}

	var cdval data.CustomTime
	cdval.UnmarshalText([]byte("2015-05-21T11:30:30"))
	if r[1].DateCreated != cdval {
		t.Errorf("Invalid DateCreated: %s  should be: %s", r[1].DateCreated, cdval)
	}

}

func TestAddReport(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestAddReport <<<<<<<<<<<<<<<<<<<<<<<<<<")
	newRpt := data.BaseReport{
		DateCreated:       data.NewCustomTime("2015-02-20T13:45:30"),
		DateUpdated:       data.NewCustomTime("2015-02-25T09:00:01.000"),
		DeviceType:        "IPHONE",
		DeviceModel:       "5S",
		DeviceID:          "3333",
		RequestType:       "Graffiti Removal",
		RequestTypeID:     "10",
		ImageUrl:          "http://www.citysourced.com/image_200.png",
		ImageUrlXl:        "http://www.citysourced.com/image_xl_200.png",
		ImageUrlLg:        "http://www.citysourced.com/image_lg_200.png",
		ImageUrlMd:        "http://www.citysourced.com/image_md_200.png",
		ImageUrlSm:        "http://www.citysourced.com/image_sm_200.png",
		ImageUrlXs:        "http://www.citysourced.com/image_xs_200.png",
		City:              "San Jose",
		State:             "CA",
		ZipCode:           "95101",
		Latitude:          "37.339541",
		Longitude:         "-121.885229",
		Directionality:    "25 N NW",
		Description:       "Spray paint art at Horace Mann school.",
		AuthorNameFirst:   "Sylvester T.",
		AuthorNameLast:    "Cat",
		AuthorEmail:       "",
		AuthorTelephone:   "",
		AuthorIsAnonymous: "true",
		UrlDetail:         "http://www.citysourced.com/report/100/graffiti",
		UrlShortened:      "",
		StatusType:        "Open",
	}
	data.D.Append(newRpt)
	fmt.Printf("------ After add:\n%s\n", data.D.Display())
}

func TestAddReport2(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestAddReport <<<<<<<<<<<<<<<<<<<<<<<<<<")
	newRpt := data.BaseReport{
		DateCreated:       data.NewCustomTime("2015-02-20T13:45:30"),
		DateUpdated:       data.NewCustomTime("2015-02-25T09:00:01.000"),
		DeviceType:        "IPHONE",
		DeviceModel:       "5S",
		DeviceID:          "1111",
		RequestType:       "Graffiti Removal",
		RequestTypeID:     "10",
		ImageUrl:          "http://www.citysourced.com/image_200.png",
		ImageUrlXl:        "http://www.citysourced.com/image_xl_200.png",
		ImageUrlLg:        "http://www.citysourced.com/image_lg_200.png",
		ImageUrlMd:        "http://www.citysourced.com/image_md_200.png",
		ImageUrlSm:        "http://www.citysourced.com/image_sm_200.png",
		ImageUrlXs:        "http://www.citysourced.com/image_xs_200.png",
		City:              "San Jose",
		State:             "CA",
		ZipCode:           "95101",
		Latitude:          "37.336240",
		Longitude:         "-121.885862",
		Directionality:    "25 N NW",
		Description:       "New graffiti request - someone painted flames on Flames Restaurant",
		AuthorNameFirst:   "Wiley",
		AuthorNameLast:    "Coyote",
		AuthorEmail:       "",
		AuthorTelephone:   "",
		AuthorIsAnonymous: "true",
		UrlDetail:         "http://www.citysourced.com/report/100/graffiti",
		UrlShortened:      "",
		StatusType:        "Open",
	}
	data.D.Append(newRpt)
	fmt.Printf("------ After add:\n%s\n", data.D.Display())
}

func TestAddComment1(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestAddComment1 <<<<<<<<<<<<<<<<<<<<<<<<<<")
	if e := data.NewComment(102, data.NewCustomTime("2015-11-08T08:08:08"), "Comment created by test case"); e != nil {
		t.Errorf("NewComment() failed - error: %q", e)
	}
	fmt.Printf("------ After add:\n%s\n", data.DisplayCommentData())
}

func TestFindDeviceId(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestFindDeviceId <<<<<<<<<<<<<<<<<<<<<<<<<<")
	di := "1111"
	reports := data.D

	rpts, err := reports.FindDeviceID(di)
	if err != nil {
		t.Errorf("FindDeviceId failed - error: %q", err)
	}
	fmt.Printf("Reports found for device ID %q:\n%s", di, spew.Sdump(rpts))

	di = "2222"
	rpts, err = reports.FindDeviceID(di)
	if err != nil {
		t.Errorf("FindDeviceId failed - error: %q", err)
	}
	fmt.Printf("Reports found for device ID %q:\n%s", di, spew.Sdump(rpts))

}

func TestFindAddress(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestFindAddress <<<<<<<<<<<<<<<<<<<<<<<<<<")

	// addr := "200 E. Santa Clara St, San Jose, CA"
	addr := "73 N 6th St., San Jose, CA"
	radius := 500.0
	reports := data.D
	limit := int64(2)

	rpts, err := reports.FindAddress(addr, radius, limit)
	if err != nil {
		t.Errorf("FindAddress failed - error: %q", err)
	}
	fmt.Printf("Reports found for address %q, radius: %v:\n%s", addr, radius, spew.Sdump(rpts))

}

func TestDistance(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestDistance <<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Printf("------ Start TestDistance:\n%s\n", data.D.Display())
	rlat := 37.151079
	rlon := -121.602551
	dist := 0.0
	var (
		dvals [5]float64
	)

	dvals[0] = 1788.1925774420406
	dvals[1] = 32654.521037160826
	dvals[2] = 24778.639830370197
	dvals[3] = 32674.419251059397
	dvals[4] = 32483.464206297744

	for i, r := range data.D.Reports {
		fmt.Printf("-- i: %d  lat: %v  lng: %v\n", i, r.LatitudeV, r.LongitudeV)
		dist = r.CalcDistance(rlat, rlon)
		fmt.Printf("ID: %v at %v:%v - distance: %v\n", r.ID, r.LatitudeV, r.LongitudeV, dist)
		if i < len(dvals) {
			if dist != dvals[i] {
				t.Errorf("Wrong distance: %v  should be: %v", dist, dvals[i])
			}
		} else {
			t.Errorf("Missing test value for Reports[%d]", i)
		}
	}

}

func TestRepeatReadData(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestRepeatReadData <<<<<<<<<<<<<<<<<<<<<<<<<<")
	if err := data.Init("../data.json"); err == nil {
		t.Errorf("Duplicate calls to data.Init() should have resulted in a warning")
	}
	fmt.Println("   (Should have just received a WARN error 'Duplicate calls...')")

}
