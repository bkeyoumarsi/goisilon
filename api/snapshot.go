package goisilon

import (
	"bytes"
	"encoding/json"
	"errors"

	types "github.com/bkeyoumarsi/goisilon/types/v1"
	"github.com/mitchellh/mapstructure"
)

func (c *IsiClient) CreateSnapshot(req types.SnapshotCreate) (string, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return "", errors.New("Failed to encode request to json")
	}

	resp, err := c.HttpClient.Do("POST", Snapshots, nil, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	var snapID types.SnapshotID
	err = mapstructure.Decode(resp, &snapID)
	if err != nil {
		return "", errors.New("Failed to get snap-id from response")
	}

	return snapID.Id, nil
}
