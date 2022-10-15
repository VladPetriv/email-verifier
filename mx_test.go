package emailverifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckMx(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:  "Email with mx",
			input: "github.com",
		},
		{
			name:    "Email with mx",
			input:   "githubexists.com",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := verifier.CheckMX(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.True(t, got.HasMXRecord)
			}
		})
	}
}
