package data_test

import (
	"CitySourcedAPI/data"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestReadConfig(t *testing.T) {

	if err := data.ReadConfig("../config.json", "../data.json"); err != nil {
		t.Errorf("Error %v occurred when reading the config - data.ReadConfig()", err)
	}
	fmt.Printf("%v", data.Config.Display())
	fmt.Printf("%v", data.Data.Display())
	if data.Config.Loaded != true {
		t.Errorf("System configuration is not marked as loaded.")
	}

	ac := "1234567890"
	if a := data.Auth(ac); !a {
		t.Errorf("Auth() failed.")
	}

	ac = "1111"
	if a := data.Auth(ac); a {
		t.Errorf("Auth() passed erroneously for: %q", ac)
	}

}

func TestFindDeviceId(t *testing.T) {

	di := "123456789"
	reports := data.Data

	rpts, err := reports.FindDeviceId(di)
	if err != nil {
		t.Errorf("FindDeviceId failed - error: %q", err)
	}
	fmt.Printf("Reports found for device ID %q:\n%s", di, spew.Sdump(rpts))

	di = "987654321"
	rpts, err = reports.FindDeviceId(di)
	if err != nil {
		t.Errorf("FindDeviceId failed - error: %q", err)
	}
	fmt.Printf("Reports found for device ID %q:\n%s", di, spew.Sdump(rpts))

}

func TestDistance(t *testing.T) {
	var (
		rlat float64 = 37.151079
		rlon float64 = -121.602551
		dist float64 = 0.0
	)

	for _, r := range data.Data.Reports {
		dist = r.Distance(rlat, rlon)
		fmt.Printf("ID: %v at %v:%v - distance: %v\n", r.Id, r.Latitude, r.Longitude, dist)
	}

}

// func TestRepeatReadConfig(t *testing.T) {

// 	if err := data.ReadConfig("../config.json"); err != nil {
// 		t.Errorf("Error \"%v\" occurred when reading the config - data.ReadConfig()", err)
// 	}

// }
