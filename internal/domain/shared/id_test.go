package shared_test

import (
	"testing"

	"github.com/Yuji-Momotani/library-app-ca/internal/domain/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewID(t *testing.T) {
	t.Run("正常系: 引数のIDがそのまま返る", func(t *testing.T) {
		tests := struct {
			v        string
			expected shared.ID
		}{
			v:        "01HXYZ1234567890ABCDEFGHJK",
			expected: shared.ID("01HXYZ1234567890ABCDEFGHJK"),
		}

		actual, err := shared.NewID(tests.v)
		require.NoError(t, err)
		assert.Equal(t, tests.expected, actual)
	})

	t.Run("正常系: 引数なしの場合、新規作成", func(t *testing.T) {
		tests := struct {
			v       string
			wantErr bool
		}{
			v:       "",
			wantErr: false,
		}

		actual, err := shared.NewID(tests.v)
		assert.NoError(t, err)
		assert.NotEmpty(t, actual)
	})
}
