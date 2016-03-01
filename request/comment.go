package request

import (
	"CitySourcedAPI/data"
	"CitySourcedAPI/logs"
	"CitySourcedAPI/response"

	"fmt"
	"time"
)

// ==============================================================================================================================
//                                      CreateReportComment
// ==============================================================================================================================

type CreateReportComment struct {
	Request
	Processor
	ReportID string `xml:"ReportId" json:"ReportId"`
	reportID int64
	Comment  string `xml:"Comment" json:"Comment"`
}

func (st *CreateReportComment) Validate(start time.Time) string {
	var v validate
	st.start = start
	st.reportID = v.int("ReportID", st.ReportID)
	return v.errmsg
}

func (st *CreateReportComment) Run() (string, error) {
	err := data.NewComment(st.reportID, data.CustomTime{time.Now()}, st.Comment)
	if err != nil {
		return response.StatusMsg(fmt.Sprintf("CreateReportComment failed: %q", err), st.start), nil
	}
	return response.StatusMsg("Comment created.", st.start), nil
}

func (st CreateReportComment) String() string {
	ls := new(logs.LogString)
	ls.AddS("CreateReportComment\n")
	ls.AddS(st.Request.String())
	ls.AddF("ID %s/%d\n", st.ReportID, st.reportID)
	ls.AddF("Comment: %v\n", st.Comment)
	return ls.Box(90)
}
