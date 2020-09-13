package nullcode

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

// type Endpoint struct {
// 	Add    endpoint.Endpoint
// 	Get    endpoint.Endpoint
// 	GetAll endpoint.Endpoint
// }

func MakeAddEndpoint(svc Service) endpoint.Endpoint {
	fmt.Println("MakeAddEndpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("MakeAddEndpoint > func")

		addCodeRequest := request.(CodeRequest)
		fmt.Println("addCodeRequest", addCodeRequest)

		code, err := svc.Add(ctx)
		fmt.Println("code", code)
		if err != nil {
			return CodeResponse{error: err.Error()}, nil
		}
		return CodeResponse{id: "id"}, nil
	}
}
