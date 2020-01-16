// +build integration

package auth

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMethodMatch(t *testing.T) {
	rv, err := NewVerifier(ExcludeMethods("/sms/PublicSMS/.+", "/.+/Internal.+/.+"))
	require.NoError(t, err)
	r, ok := rv.(*verifier)
	require.True(t, ok)

	td := []struct {
		method string
		res    bool
	}{
		{
			"/sms/PublicSMS/Send",
			true,
		},
		{
			"/sms/PublicSMS/Check",
			true,
		},
		{
			"/totp/InternalTOTP/Status",
			true,
		},
		{
			"/totp/InternalTOTP/Check",
			true,
		},
	}

	for _, d := range td {
		t.Run(d.method, func(t *testing.T) {
			assert.Equal(t, d.res, r.matchMethod(d.method))
		})
	}
}
