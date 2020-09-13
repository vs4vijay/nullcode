package nullcode

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type CodeRequest struct {
	code string
}

type CodeResponse struct {
	id      string
	error   string
	message string
}

func DecodeAddRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("DecodeAddRequest")

	var request CodeRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	fmt.Println("DecodeAddRequest > request", request)

	return request, nil
}

func EncodeAddResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println("EncodeAddResponse")

	return json.NewEncoder(w).Encode(response)
}
