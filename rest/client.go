package rest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	types "github.com/bkeyoumarsi/goisilon/types/v1"
)

type Client struct {
	httpClient *http.Client
	endpoint   string
	usr        string
	pass       string
}

func NewClient(endpoint, usr, pass string, insecure bool) *Client {
	var httpClient *http.Client
	if insecure {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient = &http.Client{Transport: tr}
	} else {
		httpClient = http.DefaultClient
	}
	return &Client{httpClient, endpoint, usr, pass}
}

func (r *Client) Get(api string) (interface{}, error) {
	url := fmt.Sprintf("%s%s", r.endpoint, api)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if err := responseCheck(resp); err != nil {
		return nil, err
	}

	return getJsonBody(resp)
}

func getJsonBody(resp *http.Response) (interface{}, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func responseCheck(resp *http.Response) error {
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		var errMessages types.Errors
		err := json.Unmarshal(body, &errMessages)
		if err != nil {
			return err
		}

		var buffer bytes.Buffer
		for _, e := range errMessages.Errors {
			buffer.WriteString(e.Message)
			buffer.WriteString(" ")
		}
		return errors.New(buffer.String())
	}
	return nil
}
