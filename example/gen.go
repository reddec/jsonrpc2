package example

type Profile struct {
}

type User interface {
	Profile(token string) (*Profile, error)
	privateSum(a, b int) (int, error)
}
