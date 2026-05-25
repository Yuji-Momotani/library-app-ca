package bookdm_test

import (
	"testing"

	"github.com/Yuji-Momotani/library-app-ca/internal/domain/bookdm"
	"github.com/stretchr/testify/assert"
)

func TestNewISBN(t *testing.T) {
	tests := []struct {
		name     string
		v        string
		expected string
		wantErr  bool
	}{
		{
			name:     "正常系: ハイフンなし",
			v:        "9783161484100",
			expected: "9783161484100",
			wantErr:  false,
		},
		{
			name:     "正常系: ハイフンあり",
			v:        "978-3-16-148410-0",
			expected: "978-3-16-148410-0",
			wantErr:  false,
		},
		{
			name:    "異常系: 13桁ではない",
			v:       "12345678901234",
			wantErr: true,
		},
		{
			name:    "異常系: 数字ではない",
			v:       "1234abc890123",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isbn, err := bookdm.NewIsbn(tt.v)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Empty(t, err)
				assert.Equal(t, tt.expected, isbn.Value())
			}
		})
	}
}
