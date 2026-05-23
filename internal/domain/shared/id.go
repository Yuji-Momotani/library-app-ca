package shared

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

type ID string

func NewID(v string) (ID, error) {
	if v != "" {
		return ID(v), nil
	}

	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, rand.Reader)
	if err != nil {
		return "", err
	}
	fmt.Print(id)

	return ID(id.String()), nil
}

func (v ID) String() string {
	return string(v)
}

func (v ID) Equal(other ID) bool {
	return v == other
}
