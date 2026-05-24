package bookdm_test

import (
	"crypto/rand"
	"testing"
	"time"

	"github.com/Yuji-Momotani/library-app-ca/internal/domain/bookdm"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
)

func Test_NewBookID(t *testing.T) {
	t.Run("NewID: 正常系", func(t *testing.T) {
		id, err := bookdm.NewBookID()
		assert.NoError(t, err)
		assert.NotEmpty(t, id)
	})
}

func Test_BookIDFromString(t *testing.T) {
	ms := ulid.Timestamp(time.Now())
	id, _ := ulid.New(ms, rand.Reader)
	tests := []struct {
		name    string
		v       string
		wantErr bool
	}{
		{
			name:    "正常系: BookIDFromString",
			v:       id.String(),
			wantErr: false,
		},
		{
			name:    "異常系: ulid以外の文字列が渡された場合",
			v:       "test",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := bookdm.BookIDFromString(tt.v)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NotEmpty(t, id)
			}
		})
	}
}
