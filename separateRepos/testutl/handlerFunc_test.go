package testutl

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-starter/api/zoo"
)

func TestHandlerFuncBodyFails(t *testing.T) {
	tests := []struct {
		name   string
		method string
		url    string
	}{
		// http.validMethod(string) accepts any CHAR except CTLs or separators as a valid HTTP Method.
		{"invalid HTTP method", string([]byte{0, 1}), "/"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Create a new testing.T variable so the expected assertion failures within HandlerFunc don't incorrectly fail the tests.
				tst := &testing.T{}
				rr, body := HandlerFuncBody(tst, zoo.Status, tt.method, tt.url, "", -1)
				assert.Nil(t, rr)
				assert.Empty(t, body)
			},
		)
	}
}

func TestHandlerFuncBodySuccessful(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		url        string
		statusCode int
	}{
		{"empty method is GET", "", "/", http.StatusOK},
		{"any alphanumeric method is ok", "zzzzzzzzz", "/", http.StatusOK},
		{"post method works", http.MethodPost, "booya", http.StatusOK},
		{"empty url works", http.MethodPost, "", http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				rr, body := HandlerFuncBody(t, zoo.Status, tt.method, tt.url, "", tt.statusCode)
				assert.NotNil(t, rr)
				assert.NotEmpty(t, body)
			},
		)
	}
}

func TestHandlerFuncBodyUrlFormat(t *testing.T) {
	url := fmt.Sprintf("/%s/%s/%s", "make", "model", "series")
	rr, body := HandlerFuncBody(t, zoo.Status, http.MethodConnect, url, "", http.StatusOK)
	assert.NotNil(t, rr)
	assert.NotEmpty(t, body)
}

func TestHandlerFuncPanics(t *testing.T) {
	assert.Panics(
		t, func() {
			HandlerFunc(t, nil, http.MethodGet, "/", "", 0)
		},
	)
}
