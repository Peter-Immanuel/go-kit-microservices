package main

import (
	"errors"
	"strings"
)

// Define your Requirements or Globally used variables
var ErrEmpty = errors.New("empty String")

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
