package app

import "context"

func page(ctx context.Context) map[string]interface{} {
	return map[string]interface{}{
		"Flash": getSession(ctx).Flash().Values(),
	}
}
