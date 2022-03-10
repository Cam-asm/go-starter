package zuid_test

import (
	"errors"
	"testing"
	"testing/iotest"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	u, err := New()
	require.NoError(t, err)
	require.Regexp(t, `[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}$`, u.String())
}

// TestNewRetries illustrates how to test failures when generating UUIDs.
func TestNewRetries(t *testing.T) {
	errExpected := errors.New("new read error")
	// Change the io.Reader from crypto to one that will always return an error.
	uuid.SetRand(iotest.ErrReader(errExpected))

	u, err := New()
	require.Equal(t, errExpected, err)
	require.Equal(t, uuid.UUID{}, u)
	require.Equal(t, "00000000-0000-0000-0000-000000000000", u.String())
}
