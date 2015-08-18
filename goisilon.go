package goisilon

import (
	"errors"
	"net/url"
	"os"
	"strconv"

	isiApi "github.com/bkeyoumarsi/goisilon/api"
)

func New() (*isiApi.IsiClient, error) {
	endpoint := os.Getenv("GOISILON_ENDPOINT")
	if endpoint != "" {
		_, err := url.ParseRequestURI(endpoint)
		if err != nil {
			return nil, errors.New("GOISILON_ENDPOINT: invalid endpoint address")
		}
	} else {
		return nil, errors.New("GOISILON_ENDPOINT: missing value")
	}

	insecure, err := strconv.ParseBool(os.Getenv("GOISILON_INSECURE"))
	if err != nil {
		return nil, errors.New("GOISILON_INSECURE: invalid argument")
	}

	username := os.Getenv("GOISILON_USERNAME")
	if username == "" {
		return nil, errors.New("GOISILON_USERNAME: missing value")
	}

	password := os.Getenv("GOISILON_PASSWORD")
	if password == "" {
		return nil, errors.New("GOISILON_PASSWORD: missing value")
	}

	return isiApi.NewApiClient(endpoint, username, password, insecure), nil
}
