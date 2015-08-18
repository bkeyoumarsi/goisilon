package goisilon

import (
	"testing"

	types "github.com/bkeyoumarsi/goisilon/types/v1"
)

const (
	endpoint = "https://10.28.102.200:8080"
	username = "root"
	password = "a"
	insecure = true
)

func TestCreateDirectory(t *testing.T) {
	headers := make(map[string]string)
	headers["x-isi-ifs-access-control"] = "public_read_write"

	c := NewApiClient(endpoint, username, password, insecure)
	err := c.CreateDirectory("/ifs/data/temp_dir", headers, false)
	if err != nil {
		t.Fail()
	}
}

func TestDirectoryExists(t *testing.T) {
	c := NewApiClient(endpoint, username, password, insecure)
	exists, err := c.DirectoryExists("/ifs/data/", "temp_dir")
	if err != nil || !exists {
		t.Fail()
	}
}

func TestUpdateAcl(t *testing.T) {
	c := NewApiClient(endpoint, username, password, insecure)
	var req = types.AclRequest{
		Authoritative: "acl",
		Action:        "update",
		Owner:         types.Ownership{Id: "UID:65534", Name: "nobody", Type: "user"},
		Group:         types.Ownership{Id: "GID:65534", Name: "nobody", Type: "group"},
	}
	err := c.UpdateDirAcl("/ifs/data/temp_dir", req)
	if err != nil {
		t.Fail()
	}
}

func TestDeleteDirectory(t *testing.T) {
	c := NewApiClient(endpoint, username, password, insecure)
	err := c.DeleteDirectory("/ifs/data/temp_dir", false)
	if err != nil {
		t.Fail()
	}
	exists, err := c.DirectoryExists("/ifs/data/", "temp_dir")
	if err != nil || exists {
		t.Fail()
	}
}
