package eds

import "net/http"

type Connection struct {
	Client    *http.Client
	AuthToken string
}

func NewConnection() Connection {
	c := &http.Client{}
	return Connection{Client: c}
}
