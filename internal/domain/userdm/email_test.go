package userdm_test

import (
	"testing"

	"github.com/Yuji-Momotani/library-app-ca/internal/domain/userdm"
	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		name        string
		v           string
		expected    string
		wantErr     bool
		expectedErr error
	}{
		{
			name:     "正常系: 正しいメールアドレスフォーマット",
			v:        "test@example.com",
			expected: "test@example.com",
			wantErr:  false,
		},
		{
			name:        "異常系: メールアドレスフォーマットが正しくない",
			v:           "test1234",
			wantErr:     true,
			expectedErr: userdm.ErrInvalidEmailFormat,
		},
		{
			name:        "異常系: @で終わる",
			v:           "test1234@",
			wantErr:     true,
			expectedErr: userdm.ErrInvalidEmailFormat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := userdm.NewEmail(tt.v)
			if tt.wantErr {
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, actual.Value())
			}
		})
	}
}

func TestDomain(t *testing.T) {
	tests := []struct {
		name     string
		v        string
		expected string
	}{
		{
			name:     "ドメイン取得チェック",
			v:        "test@example.com",
			expected: "example.com",
		},
		{
			name:     "ドメイン取得チェック",
			v:        "test@hoge.co.jp",
			expected: "hoge.co.jp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email, _ := userdm.NewEmail(tt.v)
			assert.Equal(t, tt.expected, email.Domain())
		})
	}
}
