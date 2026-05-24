package bookdm

import (
	"fmt"

	"github.com/Yuji-Momotani/library-app-ca/internal/domain/shared"
	"github.com/oklog/ulid/v2"
)

type BookID shared.ID

func NewBookID() (BookID, error) {
	id, err := shared.NewID()
	if err != nil {
		return BookID(""), fmt.Errorf("failed to shared newid:%w", err)
	}
	return BookID(id), nil
}

func BookIDFromString(s string) (BookID, error) {
	ulid, err := ulid.Parse(s)
	if err != nil {
		return BookID(""), err
	}

	return BookID(shared.IDFromULID(ulid)), nil
}

func (id BookID) Value() string {
	return shared.ID(id).String()
}

func (id BookID) Equal(v BookID) bool {
	return shared.ID(id).Equal(shared.ID(v))
}
