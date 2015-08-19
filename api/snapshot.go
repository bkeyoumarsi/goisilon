package goisilon

import (
	"bytes"
	"encoding/json"
	"errors"

	types "github.com/bkeyoumarsi/goisilon/types/v1"
)

func (c *IsiClient) CreateSnapshot(req types.SnapshotCreateReq) (string, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return "", errors.New("Failed to encode request to json")
	}

	_, err = c.HttpClient.Do("POST", Snapshots, nil, bytes.NewBuffer(body), false)
	if err != nil {
		return "", err
	}

	/* Ignoring the response until PAPI handler is fixed and matches spec
	var snapID types.SnapshotID
	err = mapstructure.Decode(resp, &snapID)
	if err != nil {
		return "", errors.New("Failed to get snap-id from response")
	}

	return snapID.Id, nil
	*/
	return "", nil
}
