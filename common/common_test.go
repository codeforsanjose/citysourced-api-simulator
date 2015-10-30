package common

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestBasics(t *testing.T) {

	fmt.Printf("\n\n======================================================== TestBasics ========================================================\n")
	fmt.Println(spew.Sdump(SysInfo))

	n := SysInfo.SystemName(0)
	if n != "?" {
		t.Errorf("Invalid system name %q for id: 0", n)
	}

}

func TestShip(t *testing.T) {

	fmt.Printf("\n\n======================================================== TestShip ========================================================\n")
	ts := new(UnixTimestamp_type)
	if ts.String() != "***" {
		t.Errorf("Timestamp invalid!")
	}
	fmt.Printf("Value after initialization: %v\n", ts)

	ts.SetCurrentTime()
	fmt.Printf("Value after setting current time: %v\n", ts)

	fmt.Printf("String value: %q\n", ts.String())

	pkg := NewShipPackage()
	fmt.Printf("           **** New Ship Package ****%s", pkg)

	var fldInt int
	fldInt = 101
	pkg.Results = append(pkg.Results, ShipDataType{"intField", &fldInt})

	var fldInt32 int32
	fldInt32 = 320032
	pkg.Results = append(pkg.Results, ShipDataType{"int32Field", &fldInt32})

	var fldInt64 int64
	fldInt64 = 640064
	pkg.Results = append(pkg.Results, ShipDataType{"Int64Field", &fldInt64})

	fldStr := "Some string value"
	pkg.Results = append(pkg.Results, ShipDataType{"StringField", &fldStr})

	var fldFloat32 float32
	fldFloat32 = 32.32
	pkg.Results = append(pkg.Results, ShipDataType{"float32Field", &fldFloat32})

	var fldFloat64 float64
	fldFloat64 = 64.64
	pkg.Results = append(pkg.Results, ShipDataType{"Float64Field", &fldFloat64})

	fmt.Printf("           **** Loaded Ship Package ****%s", pkg)

}
