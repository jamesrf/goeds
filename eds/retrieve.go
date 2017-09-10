package eds

import "encoding/json"

type RetrieveRequestMessage struct {
	EbookPreferredFormat string
	HighlightTerms       []string
	An                   string
	DbId                 string
}

type RetrieveResponseMessage struct {
	Record []Record
}

const EDS_RETRIEVE_URL = "https://eds-api.ebscohost.com/EDSAPI/rest/Retrieve"

func (ses *Session) Retrieve(req RetrieveRequestMessage) (rrm *RetrieveResponseMessage, err error) {
	body, err := json.Marshal(req)
	if err != nil {
		return
	}
	resp, err := ses.newJSONReq(EDS_RETRIEVE_URL, body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&rrm)

	return

}
