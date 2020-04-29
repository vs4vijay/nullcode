package nullcode

import (
	"context"
	"encoding/json"
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
	var request CodeRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeAddResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
