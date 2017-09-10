package eds

import (
	"encoding/json"
)

const edsSearchURL = "https://eds-api.ebscohost.com/edsapi/rest/Search"

type SearchRequestMessage struct {
	SearchCriteria    SearchCriteria
	RetrievalCriteria RetrievalCriteria
	Actions           []string
}

type SearchCriteria struct {
	Queries        []Query
	SearchMode     string
	IncludeFacets  string
	FacetFilters   []FacetFilter
	Limiters       []Limiter
	Expanders      []string
	Sort           string
	PublicationId  string
	RelatedContent []string
	AutoSuggest    string
	AutoCorrect    string
}

type Query struct {
	BooleanOperator string
	FieldCode       string
	Term            string
}
type Limiter struct {
	Id     string
	Values []string
}

type FacetFilter struct {
	FilterId    string
	FacetValues []FacetValue
}
type FacetValue struct {
	Id    string
	Value string
}

type RetrievalCriteria struct {
	View           string
	ResultsPerPage int32
	PageNumber     int32
	Highlight      string
}

func NewTestSearch() SearchRequestMessage {
	srm := SearchRequestMessage{}

	srm.SearchCriteria.Queries = []Query{{FieldCode: "TI", Term: "Babbler", BooleanOperator: "AND"}}
	srm.SearchCriteria.SearchMode = "all"
	srm.SearchCriteria.IncludeFacets = "y"
	srm.SearchCriteria.Sort = "relevance"
	srm.SearchCriteria.AutoCorrect = "n"
	srm.SearchCriteria.AutoSuggest = "n"
	srm.RetrievalCriteria = RetrievalCriteria{Highlight: "y",
		View: "brief", ResultsPerPage: 20, PageNumber: 1}

	return srm
}

func (ses *Session) Search(srm SearchRequestMessage) (sr SearchResponseMessage, err error) {

	body, err := json.Marshal(srm)
	if err != nil {
		return
	}

	resp, err := ses.newJSONReq(edsSearchURL, body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&sr)
	if err != nil {
		return
	}
	return

}
