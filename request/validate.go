package request

import (
	"fmt"
	"strconv"
)

// ==============================================================================================================================
//                                      Validate
// ==============================================================================================================================

type validate struct {
	errmsg string
}

func newValidate() *validate {
	return new(validate)
}

func (v *validate) float(name, val string) float64 {
	setToDflt := false
	vp, ok := vparms[name]
	if !ok {
		v.errmsg = v.errmsg + fmt.Sprintf("Attempting to validate an unkown field: %q", name)
		return 0.0
	}

	if val == "" {
		val = vp.dflt
		setToDflt = true
	}
	out, err := strconv.ParseFloat(val, 64)
	if (err != nil) || (setToDflt && vp.required) {
		v.errmsg = v.errmsg + fmt.Sprintf("Error %q converting %s: %q\n", err, name, val)
	}
	return out
}

func (v *validate) int(name, val string) int64 {
	setToDflt := false
	vp, ok := vparms[name]
	if !ok {
		v.errmsg = v.errmsg + fmt.Sprintf("Attempting to validate an unkown field: %q", name)
		return 0
	}

	if val == "" {
		val = vp.dflt
		setToDflt = true
	}
	out, err := strconv.ParseInt(val, 10, 64)
	if (err != nil) || (setToDflt && vp.required) {
		v.errmsg = v.errmsg + fmt.Sprintf("Error %q converting %s: %q\n", err, name, val)
	}
	return out
}

func (v *validate) bool(name, val string) bool {
	setToDflt := false
	vp, ok := vparms[name]
	if !ok {
		v.errmsg = v.errmsg + fmt.Sprintf("Attempting to validate an unkown field: %q", name)
		return false
	}

	if val == "" {
		val = vp.dflt
		setToDflt = true
	}
	out, err := strconv.ParseBool(val)
	if (err != nil) || (setToDflt && vp.required) {
		v.errmsg = v.errmsg + fmt.Sprintf("Error %q converting %s: %q\n", err, name, val)
	}
	return out
}

type vparm struct {
	vtype    string
	required bool
	dflt     string
}

var vparms map[string]vparm

func init() {
	vparms = make(map[string]vparm)

	vparms["ReportID"] = vparm{"int", true, "0"}
	vparms["Latitude"] = vparm{"float", true, "0.0"}
	vparms["Longitude"] = vparm{"float", true, "0.0"}
	vparms["AuthorIsAnonymous"] = vparm{"bool", false, "true"}
	vparms["Radius"] = vparm{"float", false, "100.0"}
	vparms["MaxResults"] = vparm{"int", false, "10"}
	vparms["IncludeDetails"] = vparm{"bool", false, "false"}
	vparms["IncludeComments"] = vparm{"bool", false, "false"}
	vparms["IncludeVotes"] = vparm{"bool", false, "false"}
	vparms["Votes"] = vparm{"bool", false, "0"}
}
