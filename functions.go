package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xabi93/cloud-function/math"
)

type Request struct {
	First  int `json:"first,omitempty"`
	Second int `json:"second,omitempty"`
}

type Response struct {
	Result int `json:"result,omitempty"`
}

func Sum(w http.ResponseWriter, r *http.Request) {
	var req Request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(Response{math.Sum(req.First, req.Second)})
	return
}

func Substract(w http.ResponseWriter, r *http.Request) {
	var req Request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(Response{math.Substract(req.First, req.Second)})
	return
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func SumAsync(ctx context.Context, msg PubSubMessage) error {
	var req Request
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		return err
	}

	result := math.Sum(req.First, req.Second)
	if result%2 == 0 {
		return fmt.Errorf("result: %d is even", result)
	}

	return nil
}
