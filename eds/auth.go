package eds

import (
	"bytes"
	"encoding/json"
	"log"
)

// edsUidAuthUrl
const edsUIDAuthURL = "https://eds-api.ebscohost.com/authservice/rest/UIDAuth"
const edsIPAuthURL = "https://eds-api.ebscohost.com/authservice/rest/IPAuth"

// AuthRequestMessage is a message sent to EDS to request authentication
type AuthRequestMessage struct {
	UserID   string
	Password string
}

// AuthResponseMessage is a message sent by EDS to the client with
// the result of the authentication request
type AuthResponseMessage struct {
	AuthToken   string
	AuthTimeout string
}

// AuthenticateUser authenticates a given userid/password against EDS
func (c *Connection) AuthenticateUser(userid string, password string) error {

	a := AuthRequestMessage{UserID: userid, Password: password}
	body, err := json.Marshal(a)
	if err != nil {
		return err
	}
	err = c.authRequest(body, edsUIDAuthURL)
	if err != nil {
		return err
	}
	return nil

}

// AuthenticateIP authenticates EDS via IP address
func (c *Connection) AuthenticateIP() error {
	body := []byte("{}")
	err := c.authRequest(body, edsIPAuthURL)
	if err != nil {
		return err
	}
	return nil
}

func (c *Connection) authRequest(body []byte, baseURL string) (err error) {
	log.Printf("Logging in...")
	resp, err := c.Client.Post(baseURL, "application/json", bytes.NewBuffer(body))

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var a AuthResponseMessage
	err = json.NewDecoder(resp.Body).Decode(&a)
	if err != nil {
		return err
	}
	log.Printf("Logged in.  Timeout %s\n", a.AuthTimeout)

	c.AuthToken = a.AuthToken
	return

}

// IsAuthenticated checks if the current connection is authenticated
func (c Connection) IsAuthenticated() bool {
	if c.AuthToken != "" {
		return true
	}
	return false
}
