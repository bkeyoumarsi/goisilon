package goisilon

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type DirChildren struct {
	Name string `json:"name"`
}

type Directory struct {
	Children []DirChildren `json:"children"`
}

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

	_, err := c.HttpClient.Do("PUT", url, headers, nil)
	return err
}

// Path arguemnt must start with /ifs/
func (c *IsiClient) DirectoryExists(parentDir, dirName string) (bool, error) {
	url := fmt.Sprintf("/namespace/%s", parentDir)

	data, err := c.HttpClient.Do("GET", url, nil, nil)
	if err != nil {
		return false, err
	}

	var content Directory
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

	_, err := c.HttpClient.Do("DELETE", url, nil, nil)
	return err
}
