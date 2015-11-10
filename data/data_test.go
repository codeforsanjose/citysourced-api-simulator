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
		t.Errorf("Attempting to load an invalid file should have caused an error", err)
	}
	fmt.Println("   (Should have just received a CRIT error 'Failed to open...')")
}

func TestReadDataInvalidJSON(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestReadDataInvalidJSON <<<<<<<<<<<<<<<<<<<<<<<<<<")
	// Load Data
	if err := data.Init("tests/data_faulty.json"); err == nil {
		t.Errorf("Attempting to load a faulty file should have caused an error", err)
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
func TestDataValidity(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestDataValidity <<<<<<<<<<<<<<<<<<<<<<<<<<")
	// Check lastID
	if lId := data.D.LastId(); lId != 102 {
		t.Errorf("LastId: %v is incorrect.", lId)
	}

	// Make sure we've got the data we think we should have - check random data:
	r, e := data.D.FindId(101)
	if e != nil {
		t.Errorf("FindId failed: %q.", e)
	}
	fmt.Printf("Id 100 - type: %T.\n", r)

	var cdval data.CustomTime
	cdval.UnmarshalText([]byte("2015-10-20T13:45:30"))
	if r.DateCreated != cdval {
		t.Errorf("Invalid DateCreated: %s  should be: %s", r.DateCreated, cdval)
	}
	sval := "101101101"
	if r.DeviceId != sval {
		t.Errorf("Invalid DeviceId: %s  should be: %s", r.DeviceId, sval)
	}
	var fval float64 = -121.886329
	if r.Lng() != fval {
		t.Errorf("Invalid Longitude: %v  should be: %v", r.Lng(), fval)
	}

	bval := true
	if r.AuthIsAnon() != bval {
		t.Errorf("Invalid AuthorIsAnonymous: %v  should be: %v", r.AuthIsAnon(), bval)
	}
}

func TestAddReport(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestAddReport <<<<<<<<<<<<<<<<<<<<<<<<<<")
	newRpt := data.BaseReport_Type{
		DateCreated:       data.NewCustomTime("2015-02-20T13:45:30"),
		DateUpdated:       data.NewCustomTime("2015-02-25T09:00:01.000"),
		DeviceType:        "IPHONE",
		DeviceModel:       "5S",
		DeviceId:          "new01",
		RequestType:       "Graffiti Removal",
		RequestTypeId:     "10",
		ImageUrl:          "http://www.citysourced.com/image_200.png",
		ImageUrlXl:        "http://www.citysourced.com/image_xl_200.png",
		ImageUrlLg:        "http://www.citysourced.com/image_lg_200.png",
		ImageUrlMd:        "http://www.citysourced.com/image_md_200.png",
		ImageUrlSm:        "http://www.citysourced.com/image_sm_200.png",
		ImageUrlXs:        "http://www.citysourced.com/image_xs_200.png",
		City:              "San Jose",
		State:             "CA",
		ZipCode:           "95101",
		Latitude:          "37.338208",
		Longitude:         "-121.886329",
		Directionality:    "25 N NW",
		Description:       "New graffiti request",
		AuthorNameFirst:   "Sylvester T.",
		AuthorNameLast:    "Cat",
		AuthorEmail:       "",
		AuthorTelephone:   "",
		AuthorIsAnonymous: "true",
		UrlDetail:         "http://www.citysourced.com/report/100/graffiti",
		UrlShortened:      "",
		StatusType:        "Open",
	}
	data.D.AddReport(newRpt)
	fmt.Printf("------ After add:\n%s\n", data.D.Display())
}

func TestFindDeviceId(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestFindDeviceId <<<<<<<<<<<<<<<<<<<<<<<<<<")
	di := "100102100102"
	reports := data.D

	rpts, err := reports.FindDeviceId(di)
	if err != nil {
		t.Errorf("FindDeviceId failed - error: %q", err)
	}
	fmt.Printf("Reports found for device ID %q:\n%s", di, spew.Sdump(rpts))

	di = "101101101"
	rpts, err = reports.FindDeviceId(di)
	if err != nil {
		t.Errorf("FindDeviceId failed - error: %q", err)
	}
	fmt.Printf("Reports found for device ID %q:\n%s", di, spew.Sdump(rpts))

}

func TestFindAddress(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestFindAddress <<<<<<<<<<<<<<<<<<<<<<<<<<")

	addr := "200 E. Santa Clara St, San Jose, CA"
	radius := 100.0
	reports := data.D

	rpts, err := reports.FindAddress(addr, radius)
	if err != nil {
		t.Errorf("FindAddress failed - error: %q", err)
	}
	fmt.Printf("Reports found for address %q, radius: %v:\n%s", addr, radius, spew.Sdump(rpts))

}

func TestDistance(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestDistance <<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Printf("------ Start TestDistance:\n%s\n", data.D.Display())
	var (
		rlat  float64 = 37.151079
		rlon  float64 = -121.602551
		dist  float64 = 0.0
		dvals [4]float64
	)

	dvals[0] = 1788.1925774420406
	dvals[1] = 32654.521037160826
	dvals[2] = 24778.639830370197
	dvals[3] = 32654.521037160826

	for i, r := range data.D.Reports {
		fmt.Printf("-- i: %d  lat: %v  lng: %v\n", i, r.Lat(), r.Lng())
		dist = r.Distance(rlat, rlon)
		fmt.Printf("ID: %v at %v:%v - distance: %v\n", r.Id, r.Lat(), r.Lng(), dist)
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
