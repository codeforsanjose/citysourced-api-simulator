package data

import (
	"time"
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
