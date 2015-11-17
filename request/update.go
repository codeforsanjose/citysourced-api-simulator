package request

import (
	"CitySourcedAPI/data"
	"CitySourcedAPI/logs"
	"CitySourcedAPI/response"

	"time"

	"github.com/davecgh/go-spew/spew"
)

// ==============================================================================================================================
//                                      UpdateThreeOneOne
// ==============================================================================================================================

type UpdateThreeOneOne struct {
	Request
	Processor
	ID string `xml:"TicketId" json:"TicketId"`
	id int64

	TicketSLA string `xml:"TicketSla" json:"TicketSla"`
}

func (st *UpdateThreeOneOne) Validate(start time.Time) string {
	var v validate
	st.start = start
	st.id = v.int("ID", st.ID)
	return v.errmsg
}

func (st *UpdateThreeOneOne) Run() (string, error) {
	rpts, _ := data.D.FindAddress(st.Address, st.radius, st.maxResults)
	log.Debug(">>> rpts:\n%s\n", spew.Sdump(rpts))

	resp, _ := response.NewResponseReports(true, st.Start(), rpts)
	return resp, nil
}

func (st UpdateThreeOneOne) String() string {
	ls := new(logs.LogString)
	ls.AddS("UpdateThreeOneOne\n")
	ls.AddS(st.Request.String())
	ls.AddF("Address \"%v\"\n", st.Address)
	ls.AddF("Radius %s/%v   MaxResults: %s/%v\n", st.Radius, st.radius, st.MaxResults, st.maxResults)
	ls.AddF("IncludeDetails: %v/%t\n", st.IncludeDetails, st.includeDetails)
	ls.AddF("Date Range: %v  to: %v \n", st.DateRangeStart, st.DateRangeEnd)
	return ls.Box(90)
}
