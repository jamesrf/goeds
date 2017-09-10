package eds

import "net/http"

// Connection represets an authenticated connection to EDS
type Connection struct {
	Client    *http.Client
	AuthToken string
}

// NewConnection creates a new EDS connection
func NewConnection() Connection {
	c := &http.Client{}
	return Connection{Client: c}
}
