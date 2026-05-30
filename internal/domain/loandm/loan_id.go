package loandm

import (
	"fmt"

	"github.com/Yuji-Momotani/library-app-ca/internal/domain/shared"
	"github.com/oklog/ulid/v2"
)

type LoanID shared.ID

func NewLoanID() (*LoanID, error) {
	id, err := shared.NewID()
	if err != nil {
		return nil, fmt.Errorf("failed to shared NewID: %w", err)
	}

	lID := LoanID(id)

	return &lID, nil
}

func LoanIDFromString(v string) (*LoanID, error) {
	sid, err := ulid.Parse(v)
	if err != nil {
		return nil, fmt.Errorf("failed to ulid parse:%w", err)
	}

	id := shared.IDFromULID(sid)
	LoanID := LoanID(id)

	return &LoanID, nil
}

func (v LoanID) Value() string {
	return shared.ID(v).String()
}

func (v LoanID) Equal(other LoanID) bool {
	return shared.ID(v).Equal(shared.ID(other))
}
