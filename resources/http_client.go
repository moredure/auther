package resources

import (
	"net/http"
	"time"
)

func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}