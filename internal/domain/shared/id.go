package shared

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

type ID string

func NewID() (ID, error) {
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, rand.Reader)
	if err != nil {
		return "", fmt.Errorf("failed to ulid new %w", err)
	}
	fmt.Print(id)

	return ID(id.String()), nil
}

func IDFromULID(id ulid.ULID) ID {
	return ID(id.String())
}

func (v ID) String() string {
	return string(v)
}

func (v ID) Equal(other ID) bool {
	return v == other
}
