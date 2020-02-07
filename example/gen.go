package example

import (
	bg "math/big"
	"time"
)

type Profile struct {
}

// General user profile access
type User interface {
	// Get user profile
	Profile(token string, at time.Time, val *bg.Int) (*Profile, error)
	privateSum(a, b int) (int, error)
}
