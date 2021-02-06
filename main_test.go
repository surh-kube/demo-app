package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"bitbucket.it.ittreasury.com/gitops/demo/server"
	"github.com/go-playground/assert/v2"
)

func TestPing(t *testing.T) {
	s := server.NewGinServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	s.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}
