package goisilon

import (
	"os"
	"strconv"

	isiApi "github.com/bkeyoumarsi/goisilon/api"
	"github.com/bkeyoumarsi/goisilon/rest"
)

func New() (*isiApi.IsiClient, error) {
	endpoint := os.Getenv("GOISILON_ENDPOINT")
	insecure, _ := strconv.ParseBool(os.Getenv("GOISILON_INSECURE"))
	username := os.Getenv("GOXTREMIO_USERNAME")
	password := os.Getenv("GOXTREMIO_PASSWORD")

	r := rest.NewClient(endpoint, username, password, insecure)

	return isiApi.NewApiClient(r), nil
}
