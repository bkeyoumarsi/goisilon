package goisilon

import (
	"github.com/bkeyoumarsi/goisilon/rest"
)

type IsiClient struct {
	HttpClient *rest.Client
}

func NewApiClient(endpoint, username, password string, insecure bool) *IsiClient {
	r := rest.NewClient(endpoint, username, password, insecure)
	return &IsiClient{r}
}
