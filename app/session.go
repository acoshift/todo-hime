package app

import (
	"context"

	"github.com/acoshift/session"
)

const (
	sessName = "sess"
)

func getSession(ctx context.Context) *session.Session {
	sess, err := session.Get(ctx, sessName)
	must(err)
	return sess
}
