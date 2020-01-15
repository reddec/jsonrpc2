package example

type Profile struct {
}

// General user profile access
type User interface {
	// Get user profile
	Profile(token string) (*Profile, error)
	privateSum(a, b int) (int, error)
}
