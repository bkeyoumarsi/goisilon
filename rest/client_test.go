package rest

import "testing"

const (
	endpoint = "https://10.28.102.200:8080"
	username = "root"
	password = "a"
	insecure = true
)

func TestError(t *testing.T) {
	c := NewClient(endpoint, username, password, insecure)
	_, err := c.Get("/platform/DoesNotExist", nil)
	if err == nil {
		t.Fail()
	}
}
