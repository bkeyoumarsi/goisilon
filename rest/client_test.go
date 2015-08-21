package rest

import "testing"

const (
	endpoint = "https://10.28.104.191:8080"
	username = "root"
	password = "test_pass"
	insecure = true
)

func TestError(t *testing.T) {
	c := NewClient(endpoint, username, password, insecure)
	_, err := c.Do("GET", "/platform/DoesNotExist", nil, nil, true)
	if err == nil {
		t.Fail()
	}
}

func TestDoRequest(t *testing.T) {
	c := NewClient(endpoint, username, password, insecure)
	data, err := c.Do("GET", "/platform/1/cluster/statfs", nil, nil, false)
	if err != nil {
		t.Fail()
	}
	m := data.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			if k == "f_mntfromname" && vv != "OneFS" {
				t.Fail()
			}
		}
	}
}
