package app

import (
	"context"

	"github.com/acoshift/session"
)

const (
	sessName = "sess"
)

func getSession(ctx context.Context) *session.Session {
	return session.Get(ctx, sessName)
}
