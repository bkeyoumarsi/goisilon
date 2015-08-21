package goisilon

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	types "github.com/bkeyoumarsi/goisilon/types/v1"
	"github.com/mitchellh/mapstructure"
)

func (c *IsiClient) GetSnapshots(name string) (types.SnapshotList, error) {
	var snapshots types.SnapshotList
	var url = SnapshotsApi

	if name != "" {
		url = fmt.Sprintf("%s/%s", SnapshotsApi, name)
	}

	resp, err := c.HttpClient.Do("GET", url, nil, nil, false)
	if err != nil {
		return snapshots, err
	}

	err = mapstructure.Decode(resp, &snapshots)
	if err != nil {
		return snapshots, errors.New("Failed to get list of snapshots")
	}

	return snapshots, nil
}

func (c *IsiClient) CreateSnapshot(req types.SnapshotCreateReq) (int, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return -1, errors.New("Failed to encode request to json")
	}

	resp, err := c.HttpClient.Do("POST", SnapshotsApi,
		nil, bytes.NewBuffer(body), false)
	if err != nil {
		return -1, err
	}

	var result types.Snapshot
	err = mapstructure.Decode(resp, &result)
	if err != nil {
		return -1, errors.New("Failed to get snap-id from response")
	}

	return result.Id, nil
}

func (c *IsiClient) DeleteSnapshot(name string) error {
	url := fmt.Sprintf("%s/%s", SnapshotsApi, name)
	_, err := c.HttpClient.Do("DELETE", url, nil, nil, true)
	return err
}
