package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
	default:
		t.Errorf("type is not http.handler, but is %T", v)
	}
}

func TestNoSessionLoad(t *testing.T) {
	var myH myHandler

	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
	default:
		t.Errorf("type is not http.handler, but is %T", v)
	}
}
