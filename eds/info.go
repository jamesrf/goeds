package eds

import (
	"encoding/json"
	"errors"
	"log"
)

type InfoResponseMessage struct {
	AvailableSearchCriteria  AvailableSearchCriteria
	ViewResultSettings       ViewResultSettings
	ApplicationSettings      ApplicationSettings
	APISettings              APISettings
	DetailedErrorDescription string
	ErrorDescription         string
	ErrorNumber              string
}
type AvailableSort struct {
	Id        string
	Label     string
	AddAction string
}
type AvailableSearchField struct {
	FieldCode string
	Label     string
}
type AvailableExpander struct {
	Id        string
	Label     string
	DefaultOn string
	AddAction string
}
type LimiterValue struct {
	Value         string
	AddAction     string
	LimiterValues []string
}
type AvailableLimiter struct {
	LimiterValues []LimiterValue
	DefaultOn     string
	Order         string
}
type AvailableSearchMode struct {
	Mode      string
	Label     string
	DefaultOn string
	AddAction string
}
type AvailableRelatedContent struct {
	Type      string
	Label     string
	DefaultOn string
	AddAction string
}
type AvailableDidYouMeanOption struct {
	Id        string
	Label     string
	DefaultOn string
}
type AvailableSearchCriteria struct {
	AvailableSorts             []AvailableSort
	AvailableSearchFields      []AvailableSearchField
	AvailableExpanders         []AvailableExpander
	AvailableLimiters          []AvailableLimiter
	AvailableSearchModes       []AvailableSearchMode
	AvailableRelatedContent    []AvailableRelatedContent
	AvailableDidYouMeanOptions []AvailableDidYouMeanOption
}

type ViewResultSettings struct {
	ResultsPerPage string
	ResultListView string
}
type ApplicationSettings struct {
	SessionTimeout string
}
type APISettings struct {
	MaxRecordJumpAhead string
}

const edsInfoURL = "https://eds-api.ebscohost.com/edsapi/rest/info"

func (ses *Session) GetInfo() (irm InfoResponseMessage, err error) {

	log.Printf("Requesting info...\n")
	body := []byte("")

	resp, err := ses.newJSONReq(edsInfoURL, body)
	if err != nil {
		return irm, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&irm)
	if err != nil {
		return irm, err
	}

	if irm.ErrorNumber != "" {
		return irm, errors.New(irm.DetailedErrorDescription)
	}
	return irm, err
}
