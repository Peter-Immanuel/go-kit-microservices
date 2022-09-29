package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

//---------- Request Response Messaging Patterns ----------//

// based on our business logic we'll define request response struct for each method of our interface

// Method Uppercase RPC
type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

// Method Count RPC
type countRequest struct {
	Word string `json:"word"`
}

type countResponse struct {
	Count int `json:"count"`
}

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		// type assert the request context to get the uppercase Request struct
		req := request.(uppercaseRequest)

		// Call the Uppercase method of the service.
		v, err := svc.Uppercase(req.S)

		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		// type assert request satisfies the countRequest interface
		req := request.(countRequest)

		wordLen := svc.Count(req.Word)
		return countResponse{wordLen}, nil
	}
}

//---------- Transports ----------//
func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
