package emailverifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var verifier = NewVerifier().EnableSMTPCheck()

func TestIsFreeDomain(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedResult bool
	}{
		{
			name:           "free domain",
			input:          "yahoo.com",
			expectedResult: true,
		},
		{
			name:           "not free domain",
			input:          "github.com.",
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := verifier.IsFreeDomain(tt.input)
			assert.Equal(t, tt.expectedResult, got)
		})
	}
}

func TestIsDisposableDomain(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedResult bool
	}{
		{
			name:           "disposable domain",
			input:          "dbbd8.club",
			expectedResult: true,
		},
		{
			name:           "not disposable domain",
			input:          "gmail.com",
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := verifier.IsDisposable(tt.input)
			assert.Equal(t, tt.expectedResult, got)
		})
	}
}

func TestIsRoleAccount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedResult bool
	}{
		{
			name:           "role account",
			input:          "administrator",
			expectedResult: true,
		},
		{
			name:           "not role account",
			input:          "normal_user",
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := verifier.IsRoleAccount(tt.input)
			assert.Equal(t, tt.expectedResult, got)
		})
	}
}
