package app

import (
	"database/sql"
	"time"

	"github.com/acoshift/session"
)

// Config is the app's config
type Config struct {
	BaseURL      string
	DB           *sql.DB
	Location     *time.Location
	SessionStore session.Store
}
