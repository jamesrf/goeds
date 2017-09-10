package eds

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Session struct {
	Profile    string
	Guest      bool
	Org        string
	Token      string
	Connection *Connection
}

type CreateSessionResponseMessage struct {
	SessionToken             string
	DetailedErrorDescription string
	ErrorDescription         string
	ErrorNumber              string
}
type EndSessionMessage struct {
	SessionToken string
}
type EndSessionResponse struct {
	IsSuccessful string
}

const edsCreateSessionURL = "http://eds-api.ebscohost.com/edsapi/rest/CreateSession"
const edsEndSessionURL = "http://eds-api.ebscohost.com/edsapi/rest/endsession"

func (c *Connection) CreateSession(org string, guest bool) (ses Session, err error) {

	ses = Session{Profile: "edsapi", Org: org, Guest: guest}

	body, err := json.Marshal(ses)
	if err != nil {
		return ses, err
	}

	resp, err := c.Client.Post(edsCreateSessionURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return ses, err
	}
	defer resp.Body.Close()

	var r CreateSessionResponseMessage

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return ses, err
	}
	if r.ErrorNumber != "" {
		return ses, errors.New(r.ErrorDescription)
	}

	ses.Token = r.SessionToken
	ses.Connection = c
	return ses, err
}

func (c *Connection) EndSession(ses Session) error {

	msg := EndSessionMessage{SessionToken: ses.Token}
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	b := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", edsEndSessionURL, b)
	if err != nil {
		return err
	}
	req.Header.Set("x-authenticationToken", c.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	var r EndSessionResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return err
	}
	if r.IsSuccessful != "y" {
		return errors.New("Session end not successful")
	}
	return nil
}

func (ses *Session) newJSONReq(url string, body []byte) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return resp, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-sessionToken", ses.Token)
	req.Header.Add("x-authenticationToken", ses.Connection.AuthToken)

	resp, err = ses.Connection.Client.Do(req)
	if err != nil {
		return resp, err
	}
	return resp, err
}
