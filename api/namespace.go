package goisilon

import (
	"bytes"
	"encoding/json"
	"fmt"

	types "github.com/bkeyoumarsi/goisilon/types/v1"
	"github.com/mitchellh/mapstructure"
)

// Path arguemnt must start with /ifs/
func (c *IsiClient) CreateDirectory(path string, headers map[string]string, recursive bool) error {
	var rec = "false"
	if recursive {
		rec = "true"
	}
	url := fmt.Sprintf("/namespace/%s?recursive=%s", path, rec)

	if headers == nil {
		headers = make(map[string]string)
	}
	// Required headear parameter
	headers["x-isi-ifs-target-type"] = "container"

	_, err := c.HttpClient.Do("PUT", url, headers, nil, true)
	return err
}

// parentDir arguemnt must start with /ifs/
func (c *IsiClient) DirectoryExists(parentDir, dirName string) (bool, error) {
	url := fmt.Sprintf("/namespace/%s", parentDir)

	data, err := c.HttpClient.Do("GET", url, nil, nil, false)
	if err != nil {
		return false, err
	}

	var content types.Directory
	err = mapstructure.Decode(data, &content)
	if err != nil {
		return false, err
	}
	for _, v := range content.Children {
		if v.Name == dirName {
			return true, nil
		}
	}

	return false, nil
}

// Path arguemnt must start with /ifs/
func (c *IsiClient) DeleteDirectory(path string, recursive bool) error {
	var rec = "false"
	if recursive {
		rec = "true"
	}
	url := fmt.Sprintf("/namespace/%s?recursive=%s", path, rec)

	_, err := c.HttpClient.Do("DELETE", url, nil, nil, true)
	return err
}

func (c *IsiClient) UpdateDirAcl(path string, body types.AclUpdateReq) error {
	url := fmt.Sprintf("/namespace/%s?acl", path)

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = c.HttpClient.Do("PUT", url, nil, bytes.NewBuffer(b), true)
	return err
}
