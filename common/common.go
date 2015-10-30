package common

import (
	"fmt"
	"math"
	"strings"
	"time"

	// "github.com/davecgh/go-spew/spew"
)

// ==============================================================================================================================
//                                      Custom Time Format
// ==============================================================================================================================
// CitySourced is not using standard RFC3339 time format, at least not in their examples...

type CustomTime struct {
	time.Time
}

const customDateFmt = "2006-01-02T15:04:05"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	ct.Time, err = time.Parse(customDateFmt, string(b))
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(customDateFmt)), nil
}

func (t CustomTime) MarshalText() (result []byte, err error) {
	fmted := t.Format(customDateFmt)
	return []byte(fmted), nil
}

func (t *CustomTime) UnmarshalText(text []byte) error {
	parse, err := time.Parse(customDateFmt, string(text))
	if err != nil {
		return err
	}
	*t = CustomTime{parse}
	return nil
}

var nilTime = (time.Time{}).UnixNano()

func (ct *CustomTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}

// ==============================================================================================================================
//                                      TIMING
// ==============================================================================================================================
var programStartTime = time.Now()

func ProgramElapsedTime() float64 {
	return ToFixed(time.Since(programStartTime).Seconds(), 2)
}

// ==============================================================================================================================
//                                      MATH
// ==============================================================================================================================
func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}

// ==============================================================================================================================
//                                      STRINGS
// ==============================================================================================================================
func VarNameToGo(varName, EorP string) string {
	ss := strings.Split(varName, "_")
	for i := range ss {
		if !(EorP == "private" && i == 0) {
			ss[i] = strings.Title(ss[i])
		}
	}
	newVarName := strings.Join(ss, "")
	return newVarName
}

func BoxIt(str []string, w int) string {
	out := "\n"
	for i, ln := range str {
		if i == 0 {
			x := ((w - len(ln)) / 2) - 1
			out += fmt.Sprintf("\u2554%s %s %s\n", strings.Repeat("\u2550", x), ln, strings.Repeat("\u2550", x))
		} else {
			out += fmt.Sprintf("\u2551%s\n", strings.Replace(ln, "\n", "\n\u2551", -1))
		}
	}
	out += fmt.Sprintf("\u255A%s\n\n", strings.Repeat("\u2550", w))
	return out
}

// ==============================================================================================================================
//                                      INTERFACE TYPE
// ==============================================================================================================================
func GetInterfacePtrType(interfacePtr interface{}) string {
	if _, ok := interfacePtr.(*int); ok {
		return "int"
	}
	if _, ok := interfacePtr.(*int32); ok {
		return "int32"
	}
	if _, ok := interfacePtr.(*int64); ok {
		return "int64"
	}

	if _, ok := interfacePtr.(*float32); ok {
		return "float32"
	}
	if _, ok := interfacePtr.(*float64); ok {
		return "float64"
	}

	if _, ok := interfacePtr.(*string); ok {
		return "string"
	}
	return "unknown"
}

// ==============================================================================================================================
//                                      TIMESTAMP
// ==============================================================================================================================
type UnixTimestamp_type int64

func (this *UnixTimestamp_type) SetCurrentTime() {
	*this = UnixTimestamp_type(time.Now().Unix())
}

func (this UnixTimestamp_type) String() string {
	if this == 0 {
		return "***"
	} else {
		return fmt.Sprintf("%v", time.Unix(int64(this), 0))
	}
}
