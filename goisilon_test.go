package goisilon

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	os.Setenv("GOISILON_ENDPOINT", "https://10.28.102.200:8080")
	os.Setenv("GOISILON_INSECURE", "true")
	os.Setenv("GOISILON_USERNAME", "root")
	os.Setenv("GOISILON_PASSWORD", "a")
	c, err := New()
	if err != nil {
		t.Errorf("Failed to initiate api client with err: %s", err.Error())
		t.FailNow()
	}
	err = c.CreateDirectory("/ifs/data/test", nil, false)
	if err != nil {
		t.Errorf("Failed to create a dir using the api with err: %s", err.Error())
		t.FailNow()
	}
	exist, err := c.DirectoryExists("/ifs/data", "test")
	if err != nil || exist == false {
		t.Errorf("Checking dir existance failed with err: %s", err.Error())
	}
	err = c.DeleteDirectory("/ifs/data/test", false)
	if err != nil {
		t.Errorf("Failed to delete test directory with err: %s", err.Error())
		t.FailNow()
	}
	exist, err = c.DirectoryExists("/ifs/data", "test")
	if err != nil || exist == true {
		t.Errorf("Test directory exists after calling delete with err: %s", err.Error())
	}
}
