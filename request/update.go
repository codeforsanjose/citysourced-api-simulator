package request

import (
	"CitySourcedAPI/data"
	"CitySourcedAPI/logs"
	"CitySourcedAPI/response"

	"fmt"
	"time"
)

// ==============================================================================================================================
//                                      UpdateThreeOneOne
// ==============================================================================================================================

type UpdateThreeOneOne struct {
	Request
	Processor
	ReportID  string `xml:"TicketId" json:"TicketId"`
	reportID  int64
	TicketSLA string `xml:"TicketSla" json:"TicketSla"`
}

func (st *UpdateThreeOneOne) Validate(start time.Time) string {
	var v validate
	st.start = start
	st.reportID = v.int("ReportID", st.ReportID)
	return v.errmsg
}

func (st *UpdateThreeOneOne) Run() (string, error) {
	err := data.UpdateSLA(st.reportID, st.TicketSLA)
	if err != nil {
		return response.StatusMsg(fmt.Sprintf("Update failed: %q", err), st.start), nil
	}
	return response.StatusMsg("Report updated.", st.start), nil
}

func (st UpdateThreeOneOne) String() string {
	ls := new(logs.LogString)
	ls.AddS("UpdateThreeOneOne\n")
	ls.AddS(st.Request.String())
	ls.AddF("ID %s/%d\n", st.ReportID, st.reportID)
	ls.AddF("SLA %q\n", st.TicketSLA)
	return ls.Box(90)
}
