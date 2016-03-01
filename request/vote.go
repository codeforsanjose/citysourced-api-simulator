package request

import (
	"CitySourcedAPI/data"
	"CitySourcedAPI/logs"
	"CitySourcedAPI/response"

	"fmt"
	"time"
)

// ==============================================================================================================================
//                                      CreateReportVote
// ==============================================================================================================================

type CreateReportVote struct {
	Request
	Processor
	ReportID string `xml:"ReportId" json:"ReportId"`
	reportID int64
}

func (st *CreateReportVote) Validate(start time.Time) string {
	var v validate
	st.start = start
	st.reportID = v.int("ReportID", st.ReportID)
	return v.errmsg
}

func (st *CreateReportVote) Run() (string, error) {
	err := data.Vote(st.reportID)
	if err != nil {
		return response.StatusMsg(fmt.Sprintf("CreateReportVote failed: %q", err), st.start), nil
	}
	return response.StatusMsg("Vote logged.", st.start), nil
}

func (st CreateReportVote) String() string {
	ls := new(logs.LogString)
	ls.AddS("CreateReportVote\n")
	ls.AddS(st.Request.String())
	ls.AddF("ID %s/%d\n", st.ReportID, st.reportID)
	return ls.Box(90)
}
