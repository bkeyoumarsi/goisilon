package goisilon

import (
	"testing"

	types "github.com/bkeyoumarsi/goisilon/types/v1"
)

func init() {
	if TestClient == nil {
		TestClient = NewApiClient(TestEndpoint, TestUsername,
			TestPassword, TestInsecure)
	}
}

func TestCreateDirectory(t *testing.T) {
	headers := make(map[string]string)
	headers["x-isi-ifs-access-control"] = "public_read_write"

	err := TestClient.CreateDirectory("/ifs/data/temp_dir", headers, false)
	if err != nil {
		t.Fail()
	}
}

func TestDirectoryExists(t *testing.T) {
	exists, err := TestClient.DirectoryExists("/ifs/data/", "temp_dir")
	if err != nil || !exists {
		t.Fail()
	}
}

func TestUpdateAcl(t *testing.T) {
	var req = types.AclUpdateReq{
		Authoritative: "acl",
		Action:        "update",
		Owner:         types.Ownership{Id: "UID:65534", Name: "nobody", Type: "user"},
		Group:         types.Ownership{Id: "GID:65534", Name: "nobody", Type: "group"},
	}
	err := TestClient.UpdateDirAcl("/ifs/data/temp_dir", req)
	if err != nil {
		t.Fail()
	}
}

func TestDeleteDirectory(t *testing.T) {
	err := TestClient.DeleteDirectory("/ifs/data/temp_dir", false)
	if err != nil {
		t.Fail()
	}
	exists, err := TestClient.DirectoryExists("/ifs/data/", "temp_dir")
	if err != nil || exists {
		t.Fail()
	}
}
