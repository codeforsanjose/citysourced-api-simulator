package data

import (
	"time"
)

// ==============================================================================================================================
//                                      Custom Time Format
// ==============================================================================================================================
const customDateFmt = "2006-01-02T15:04:05"

// CitySourced is not using standard RFC3339 time format, at least not in their examples...

type CustomTime struct {
	time.Time
}

func NewCustomTime(in string) CustomTime {
	nt, _ := time.Parse(customDateFmt, in)
	return CustomTime{nt}
}

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	if len(b) == 0 {
		ct.Time = *new(time.Time)
		return
	}
	ct.Time, err = time.Parse(customDateFmt, string(b))
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(customDateFmt)), nil
}

func (ct CustomTime) MarshalText() (result []byte, err error) {
	fmted := ct.Format(customDateFmt)
	return []byte(fmted), nil
}

func (ct *CustomTime) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		*ct = CustomTime{}
		return nil
	}
	parse, err := time.Parse(customDateFmt, string(text))
	if err != nil {
		return err
	}
	*ct = CustomTime{parse}
	return nil
}

var nilTime = (time.Time{}).UnixNano()

func (ct *CustomTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}
