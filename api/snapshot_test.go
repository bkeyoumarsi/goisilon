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

func TestCreateSnapshot(t *testing.T) {
	var req = types.SnapshotCreateReq{
		Name: "TestSnapshot",
		Path: "/ifs/data",
	}
	_, err := TestClient.CreateSnapshot(req)
	if err != nil {
		t.Errorf("Failed to create snapshot on /ifs/data, err: %s", err.Error())
	}
}
