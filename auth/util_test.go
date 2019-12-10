package auth

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUID(t *testing.T) {
	td := []struct {
		name string
		ctx  context.Context
		uid  string
		err  error
	}{
		{
			"empty ctx",
			context.TODO(),
			"",
			ErrInvalidContext,
		},
		{
			"no sub",
			context.WithValue(context.TODO(), "claim", jwt.MapClaims(map[string]interface{}{})),
			"",
			ErrInvalidContext,
		},
		{
			"sub int",
			context.WithValue(context.TODO(), "claim", jwt.MapClaims(map[string]interface{}{
				"sub": 1,
			})),
			"",
			ErrInvalidContext,
		},
		{
			"correct",
			context.WithValue(context.TODO(), "claim", jwt.MapClaims(map[string]interface{}{
				"sub": "test1",
			})),
			"test1",
			nil,
		},
	}

	for _, d := range td {
		t.Run(d.name, func(t *testing.T) {
			uid, err := GetUID(d.ctx)
			assert.Equal(t, d.uid, uid)
			assert.Equal(t, d.err, err)
		})
	}
}
