package main

import (
	"bytes"
	handlers "main.go/internal"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(handlers.Create))

	resp, err := http.Post(srv.URL, "text/json", bytes.NewBuffer([]byte(`{"name":"test user ","age": 2,"friends":[]}`)))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if resp.StatusCode != 201 {
		t.Log(resp.StatusCode, "received, expected 201")
		t.Fail()
	}
}
