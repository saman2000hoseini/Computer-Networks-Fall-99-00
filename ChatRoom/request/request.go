package request

type Request interface {
	GenerateRequest() (*PackedRequest, error)
}

type PackedRequest struct {
	Type        string
	RequestBody []byte
}

func New(t string, rb []byte) *PackedRequest {
	return &PackedRequest{
		Type:        t,
		RequestBody: rb,
	}
}
