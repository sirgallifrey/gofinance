package domain

import "time"

// Does session belongs on Domain or on infrastructure???

type UserSession struct {
	Token      string
	User_id    string
	Org_id     string
	Created_at time.Time
	Expires_in time.Duration
}
