package data

import (
	"CitySourcedAPI/logs"
	"_sketches/spew"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"
)

var (
	cmtData Comments
)

func NewComment(rid int64, dc CustomTime, cmt string) error {
	return cmtData.newComment(rid, dc, cmt)
}

func FindReportComments(id int64) ([]*Comment, error) {
	return cmtData.indReportID[id], nil
}

func LastCommentID() int64 {
	return cmtData.lastID
}

func DisplayCommentData() string {
	return cmtData.String()
}

func readCommentData(filePath string) (*Comments, error) {
	if cmtData.Loaded {
		msg := "Duplicate calls to load Comment Data file!"
		log.Warning(msg)
		return &cmtData, errors.New(msg)
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		msg := fmt.Sprintf("Failed to %s", err)
		log.Critical(msg)
		return nil, errors.New(msg)
	}

	err = json.Unmarshal([]byte(file), &cmtData)
	if err != nil {
		msg := fmt.Sprintf("Invalid JSON in the Data file %q: %s", filePath, err)
		log.Critical(msg)
		return nil, errors.New(msg)
	}

	// Build Indexes
	cmtData.indexAll()
	log.Debug(spew.Sdump(cmtData.indID))

	// Update Last ID
	var lastID int64
	for _, v := range cmtData.All {
		if v.ID > lastID {
			lastID = v.ID
		}
	}
	cmtData.lastID = lastID

	cmtData.Loaded = true

	log.Debug("Comments:\n%s\n", cmtData)
	return &cmtData, nil
}

// ==============================================================================================================================
//                                      COMMENTS
// ==============================================================================================================================

type Comments struct {
	Loaded      bool
	lastID      int64
	All         []*Comment `json:"comments" xml:"comments"`
	indID       map[int64]*Comment
	indReportID map[int64][]*Comment
	sync.Mutex
}

func (c *Comments) indexAll() error {
	cmtData.indID = make(map[int64]*Comment)
	cmtData.indReportID = make(map[int64][]*Comment)
	for _, cmt := range c.All {
		c.index(cmt)
	}
	return nil
}

func (cs *Comments) newComment(rid int64, dc CustomTime, cmt string) error {
	log.Debug("[AddComment] rid: %d  time: %v  cmt: %q", rid, dc, cmt)
	// ToDo: Validate rid!

	st := Comment{
		ID:          0,
		ReportID:    rid,
		DateCreated: dc,
		Comment:     cmt,
	}

	cs.Lock()

	cs.lastID++
	st.ID = cs.lastID
	cs.All = append(cs.All, &st)

	cs.indID[cs.lastID] = &st

	if _, ok := cs.indReportID[rid]; !ok {
		cs.indReportID[rid] = make([]*Comment, 0)
	}
	cs.indReportID[rid] = append(cs.indReportID[rid], &st)

	cs.Unlock()
	return nil
}

func (c *Comments) index(indc *Comment) error {
	if !(indc.ID > 0) {
		return fmt.Errorf("Attempt to index a comment that has no ID.")
	}

	rid := indc.ReportID
	log.Debug("rid: %d", rid)

	// Index: ID
	c.indID[indc.ID] = indc

	// Index: ReportID
	if _, ok := c.indReportID[rid]; !ok {
		c.indReportID[rid] = make([]*Comment, 0)
	}
	c.indReportID[rid] = append(c.indReportID[rid], indc)
	return nil
}

func (c Comments) String() string {
	ls := new(logs.LogString)
	ls.AddS("Comments\n")
	ls.AddF("Loaded: %t   lastID: %d\n", c.Loaded, c.lastID)
	for _, x := range c.All {
		ls.AddS(x.String())
	}
	ls.AddS("ID Index  (All values should match)\n")
	for k, v := range c.indID {
		ls.AddF("   %-6d  %-6d\n", k, v.ID)
	}

	ls.AddS("ReportID Index\n")
	for k, v := range c.indReportID {
		ls.AddF("   %d  [", k)
		for _, x := range v {
			ls.AddF("%d ", x.ID)
		}
		ls.AddS("]\n")
	}

	return ls.BoxC(80)
}

// ==============================================================================================================================
//                                      COMMENT
// ==============================================================================================================================

type Comment struct {
	XMLName     xml.Name   `xml:"comment" json:"comment"`
	ID          int64      `json:"Id" xml:"Id"`
	ReportID    int64      `json:"ReportID" xml:"ReportID"`
	DateCreated CustomTime `json:"DateCreated" xml:"DateCreated"`
	Comment     string     `json:"Comment" xml:"Comment"`
}

func (c Comment) String() string {
	ls := new(logs.LogString)
	ls.AddS("Comment\n")
	ls.AddF("ID: %d   ReportID: %d  Created: %v\n", c.ID, c.ReportID, c.DateCreated)
	ls.AddF("%s\n", c.Comment)
	return ls.BoxC(70)
}
