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
	if r.Longitude != fval {
		t.Errorf("Invalid Longitude: %v  should be: %v", r.Longitude, fval)
	}

	bval := true
	if r.AuthorIsAnonymous != bval {
		t.Errorf("Invalid AuthorIsAnonymous: %v  should be: %v", r.AuthorIsAnonymous, bval)
	}
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

func TestDistance(t *testing.T) {
	fmt.Println("\n\n>>>>>>>>>>>>>>>>>>> TestDistance <<<<<<<<<<<<<<<<<<<<<<<<<<")
	var (
		rlat float64 = 37.151079
		rlon float64 = -121.602551
		dist float64 = 0.0
		// dvals [3]float64 = [1788.1925774420406, 32654.521037160826, 24778.639830370197]
		dvals [3]float64
	)

	dvals[0] = 1.1111310100377383
	dvals[1] = 20.29057239138166
	dvals[2] = 15.39672821003696

	for i, r := range data.D.Reports {
		dist = r.Distance(rlat, rlon)
		fmt.Printf("ID: %v at %v:%v - distance: %v\n", r.Id, r.Latitude, r.Longitude, dist)
		if dist != dvals[i] {
			t.Errorf("Wrong distance: %v  should be: %v", dist, dvals[i])
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
