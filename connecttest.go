package connecttest

import (
	"errors"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/require"
)

// RequireError assets that the Connect error is the expected one.
func RequireError(t *testing.T, err error, code connect.Code, message string) {
	if connecterr := new(connect.Error); errors.As(err, &connecterr) {
		require.Equal(t, connecterr.Code(), code)
		require.Equal(t, connecterr.Message(), message)
	}
}
