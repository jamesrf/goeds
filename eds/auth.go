package eds

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const UID_AUTH_BASE = "https://eds-api.ebscohost.com/authservice/rest/UIDAuth"
const IP_AUTH_BASE = "https://eds-api.ebscohost.com/authservice/rest/IPAuth"

type Auth struct {
	UserId   string
	Password string
}

type AuthResponseMessage struct {
	AuthToken   string
	AuthTimeout string
}

func (c *Connection) AuthenticateUser(userid string, password string) error {

	a := Auth{UserId: userid, Password: password}
	body, err := json.Marshal(a)
	if err != nil {
		return err
	}
	err = c.authRequest(body, UID_AUTH_BASE)
	if err != nil {
		return err
	}
	return nil

}

func (c *Connection) AuthenticateIP() error {
	body := []byte("{}")
	err := c.authRequest(body, IP_AUTH_BASE)
	if err != nil {
		return err
	}
	return nil
}

func (c *Connection) authRequest(body []byte, baseURL string) (err error) {

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
	fmt.Printf("V: %v \n\n", a)
	c.AuthToken = a.AuthToken
	return

}
