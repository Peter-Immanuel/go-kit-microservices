// Goal: Build a service that provides operation on strings
package main

import (
	"context"
	"errors"
	"strings"
	"encoding/json"


	"github.com/go-kit/kit/endpoint"
)

// Define your Requirements or Globally used variables
var ErrEmpty = errors.New("Empty String")

// Define your business Logics

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

//---------- Business logic implementation ----------//

// Define your entities
type stringService struct{}

func (stringService) Uppercase(word string) (string, error) {
	if word == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(word), nil
}

func (stringService) Count(word string) int {
	return len(word)
}

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

//---------- Endpoints ----------//
// An endpoint represent a single RPC which is a method in our service interface

// adaptors converts each method of our service into an endpoint.

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

func main(){
	svc := stringService{}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		make
	)
}