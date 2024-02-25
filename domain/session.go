package domain

import "time"

// Does session belongs on Domain or on infrastructure???

type UserSession struct {
	token      string
	user_id    string
	org_id     string
	created_at time.Time
	expires_in time.Duration
}
