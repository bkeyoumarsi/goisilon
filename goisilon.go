package goisilon

import (
	"os"
	"strconv"

	isiApi "github.com/bkeyoumarsi/goisilon/api"
)

func New() (*isiApi.IsiClient, error) {
	endpoint := os.Getenv("GOISILON_ENDPOINT")
	insecure, _ := strconv.ParseBool(os.Getenv("GOISILON_INSECURE"))
	username := os.Getenv("GOXTREMIO_USERNAME")
	password := os.Getenv("GOXTREMIO_PASSWORD")

	return isiApi.NewApiClient(endpoint, username, password, insecure), nil
}
