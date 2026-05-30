package userdm

import (
	"fmt"

	"github.com/Yuji-Momotani/library-app-ca/internal/domain/shared"
	"github.com/oklog/ulid/v2"
)

type UserID shared.ID

func NewID() (*UserID, error) {
	id, err := shared.NewID()
	if err != nil {
		return nil, err
	}

	userID := UserID(id)

	return &userID, nil
}

func UserIDFromString(v string) (*UserID, error) {
	id, err := ulid.Parse(v)
	if err != nil {
		return nil, fmt.Errorf("failed to ulid parse:%w", err)
	}

	sid := shared.IDFromULID(id)
	uid := UserID(sid)

	return &uid, nil
}

func (v UserID) Value() string {
	return shared.ID(v).String()
}

func (v UserID) Equals(other UserID) bool {
	return shared.ID(v).Equal(shared.ID(other))
}
