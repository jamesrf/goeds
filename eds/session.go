package eds

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	SessionToken string
}
type EndSessionMessage struct {
	SessionToken string
}
type EndSessionResponse struct {
	IsSuccessful bool
}

const CREATE_SESSION_URL = "http://eds-api.ebscohost.com/edsapi/rest/CreateSession"
const END_SESSION_URL = "http://eds-api.ebscohost.com/edsapi/rest/endsession"

func (c *Connection) CreateSession(org string, guest bool) (ses Session, err error) {
	ses = Session{Profile: "edsapi", Org: org, Guest: guest}
	body, err := json.Marshal(ses)
	if err != nil {
		return ses, err
	}
	resp, err := c.Client.Post(CREATE_SESSION_URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return ses, err
	}
	defer resp.Body.Close()

	var r CreateSessionResponseMessage

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return ses, err
	}

	ses.Token = r.SessionToken
	ses.Connection = c
	return ses, err
}

func (c *Connection) EndSession(ses Session) {

	msg := EndSessionMessage{SessionToken: ses.Token}
	body, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", END_SESSION_URL, bytes.NewBuffer(body))
	req.Header.Set("x-authenticationToken", c.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		panic(err)
	}

	var r EndSessionResponse
	err = json.NewDecoder(resp.Body).Decode(&r)

	fmt.Printf("%v \n", r)

}
