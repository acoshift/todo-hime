package todoapp

import (
	"context"

	"github.com/moonrhythm/session"
)

const (
	sessName = "sess"
)

func getSession(ctx context.Context) *session.Session {
	sess, err := session.Get(ctx, sessName)
	if err != nil {
		panic(err)
	}

	return sess
}
