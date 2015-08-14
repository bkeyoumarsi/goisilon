package goisilon

import (
	"github.com/bkeyoumarsi/goisilon/rest"
)

type IsiClient struct {
	HttpClient *rest.Client
}

func NewApiClient(r *rest.Client) (*IsiClient, error) {
	return &IsiClient{r}, nil
}
