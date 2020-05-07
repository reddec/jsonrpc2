package example

import (
	bg "math/big"
	"time"
)

type Status int

const (
	Active  Status = 1
	Blocked Status = 2
	Removed Status = 3
)

type Region string

const (
	APAC Region = "apac"
	EU   Region = "eu"
)

type Address struct {
	Region  Region
	Country string `json:"country"`
	City    string `json:"location,omitempty"`
}

type Meta struct {
	Status  Status
	Billing *Address `json:"billing,omitempty"`
	SubMeta *Meta    `json:"sub_meta,omitempty"`
}

type Profile struct {
	Year      *bg.Int
	Time      time.Time
	Name      string
	Age       uint8
	U16       uint16
	U32       uint32
	U64       uint64
	U         uint
	I         int
	I64       int64
	I32       int32
	I16       int16
	I8        int8
	BT        byte
	Bool      bool
	F32       float32
	F64       float64
	Addresses []Address
	Meta      Meta
}

// General user profile access
type User interface {
	// Get user profile
	Profile(token string, at time.Time, val *bg.Int) (*Profile, error)
	privateSum(a, b int) (int, error)
	Latest(times []*time.Time, num int) (time.Time, error)
}
