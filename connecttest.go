package connecttest

import (
	"errors"
	"testing"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/require"
)

// RequireError assets that the Connect error is the expected one.
func RequireError(t *testing.T, err error, code connect.Code, message string) {
	require.Error(t, err)
	if connecterr := new(connect.Error); errors.As(err, &connecterr) {
		require.Equal(t, connecterr.Code(), code)
		require.Equal(t, connecterr.Message(), message)
	}
}
