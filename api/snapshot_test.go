package goisilon

import (
	"testing"

	types "github.com/bkeyoumarsi/goisilon/types/v1"
)

const testSnapshot = "TestSnapshot"

func init() {
	if TestClient == nil {
		TestClient = NewApiClient(TestEndpoint, TestUsername,
			TestPassword, TestInsecure)
	}
}

func TestCreateSnapshot(t *testing.T) {
	var req = types.SnapshotCreateReq{
		Name: testSnapshot,
		Path: "/ifs/data",
	}
	id, err := TestClient.CreateSnapshot(req)
	if err != nil {
		t.Errorf("Failed to create snapshot on /ifs/data, err: %s", err.Error())
		return
	}
	if id == -1 {
		t.Errorf("Invalid snap-id returned: %d", id)
	}
}

func TestGetSnapshots(t *testing.T) {
	snaps, err := TestClient.GetSnapshots("")
	if err != nil {
		t.Errorf("Failed to get snapshot list, err: %s", err.Error())
		return
	}

	if snaps.Total < 1 {
		t.Error("Total returned 0 which is wrong")
		return
	}

	var createdTestSnap = false
	for _, s := range snaps.Snapshots {
		if s.Name == testSnapshot {
			createdTestSnap = true
		}
	}
	if !createdTestSnap {
		t.Errorf("Could not find %s in the returned snapshot list", testSnapshot)
	}
}

func TestDeleteSnapshots(t *testing.T) {
	err := TestClient.DeleteSnapshot(testSnapshot)
	if err != nil {
		t.Errorf("Failed to delete %s, err: %s", testSnapshot, err.Error())
		return
	}

	_, err = TestClient.GetSnapshots(testSnapshot)
	if err == nil {
		t.Error("The snapshot was not deleted")
	}
}
