// +build go1.7

package httprouter

import (
	"context"
)

type paramsKey struct{}

// GetParams gets params from context.
func GetParams(ctx context.Context) Params {
	ps, _ := ctx.Value(paramsKey{}).(Params)
	return ps
}

// GetParam gets a param by name from context
func GetParam(ctx context.Context, name string) string {
	ps := GetParams(ctx)
	if ps == nil {
		return ""
	}
	return ps.ByName(name)
}
