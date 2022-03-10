package testutl

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// HandlerFunc sets up an HTTP ResponseRecorder request for testing.
// urlPattern uses fmt.Sprintf to format according to a format specifier.
// method is one of the HTTP Methods.
// requestBody is the payload to send.
// ok returns true when the ResponseRecorder succeeded and the returned HTTP status matches statusCode,
// otherwise false if any errors occurred.
func HandlerFunc(t *testing.T, handler http.HandlerFunc, method, url, requestBody string, expectedStatus int) (rr *httptest.ResponseRecorder) {
	req, err := http.NewRequest(method, url, strings.NewReader(requestBody))
	require.NoError(t, err)

	rr = httptest.NewRecorder()

	// Make the handler function satisfy http.Handler.
	handler.ServeHTTP(rr, req)
	require.Equal(t, expectedStatus, rr.Result().StatusCode)
	return rr
}

// HandlerFuncBody is identical to HandlerFunc, but also returns the HTTP response body as a string.
func HandlerFuncBody(t *testing.T, handler http.HandlerFunc, method, urlPattern, requestBody string, expectedStatus int) (rr *httptest.ResponseRecorder, responseBody string) {
	rr = HandlerFunc(t, handler, method, urlPattern, requestBody, expectedStatus)
	if rr != nil && rr.Body != nil {
		return rr, rr.Body.String()
	}

	return rr, ""
}
