package goisilon

import (
	"os"
	"strconv"
)

func New() error {
	endpoint := os.Getenv("GOISILON_ENDPOINT")
	insecure, _ := strconv.ParseBool(os.Getenv("GOISILON_INSECURE"))
	username := os.Getenv("GOXTREMIO_USERNAME")
	password := os.Getenv("GOXTREMIO_PASSWORD")

	return nil
}
