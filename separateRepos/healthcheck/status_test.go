package healthcheck_test

import (
	"go-starter/separateRepos/healthcheck"
	"go-starter/separateRepos/testutl"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	assert.Equal(t, "/status", healthcheck.URLPattern)
}

func TestHandler(t *testing.T) {
	rr, _ := testutl.HandlerFunc(t, healthcheck.Status, http.MethodGet, healthcheck.URLPattern, nil, http.StatusOK)

	assert.Equal(t, `{"status":"healthy"}`, rr.Body.String())
}
