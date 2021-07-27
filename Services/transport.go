package Services

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

func HelloDecodeRequest(c context.Context, request *http.Request) (interface{}, error) {
	name := request.URL.Query().Get("name")
	if name == "" {
		return nil, errors.New("无效参数")
	}

	return HelloRequest{Name: name}, nil
}

func HelloEncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}
