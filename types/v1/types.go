package goisilon

type ErrObjects struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type Errors struct {
	Errors []ErrObjects `json:"errors"`
}
