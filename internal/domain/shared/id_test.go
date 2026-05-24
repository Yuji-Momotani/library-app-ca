package shared_test

import (
	"testing"

	"github.com/Yuji-Momotani/library-app-ca/internal/domain/shared"
	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	t.Run("正常系:", func(t *testing.T) {
		actual, err := shared.NewID()
		assert.NoError(t, err)
		assert.NotEmpty(t, actual)
	})
}
