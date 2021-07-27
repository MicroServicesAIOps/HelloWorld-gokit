package Services

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Reply string `json:"reply"`
}

func MakeServerEndPointHello(s IMyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(HelloRequest)

		return HelloResponse{Reply: s.Hello(r.Name)}, nil
	}
}
