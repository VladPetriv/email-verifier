package emailverifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGravatar(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            string
		expectedResult   *Gravatar
		shouldBePositive bool
	}{
		{
			name:             "Email with gravatar",
			input:            "alex@pagerduty.com",
			expectedResult:   &Gravatar{HasGravatar: true, GravatarUrl: "not empty"},
			shouldBePositive: true,
		},
		{
			name:           "Email without gravatar",
			input:          "MyemailaddressHasNoGravatar@example.com",
			expectedResult: &Gravatar{HasGravatar: false},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := verifier.CheckGravatar(tt.input)
			assert.NoError(t, err)

			if tt.shouldBePositive {
				assert.True(t, got.HasGravatar)
				assert.NotEmpty(t, got.GravatarUrl)
			} else {
				assert.False(t, got.HasGravatar)
				assert.Empty(t, got.GravatarUrl)
			}
		})
	}
}
