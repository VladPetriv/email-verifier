package emailverifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuggeestDomain(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedDomain string
	}{
		{
			name:           "DomainOK_HitExactDomain",
			input:          "gmail.com",
			expectedDomain: "",
		},
		{
			name:           "DomainOK_NullString",
			input:          "",
			expectedDomain: "",
		},
		{
			name:           "DomainOK_SimilarDomain1",
			input:          "gmaii.com",
			expectedDomain: "gmail.com",
		},
		{
			name:           "DomainOK_SimilarDomain2",
			input:          "gmai.com",
			expectedDomain: "gmail.com",
		},
		{
			name:           "DomainOK_TLD",
			input:          "gmail.edd",
			expectedDomain: "gmail.edu",
		},
		{
			name:           "DomainOK_SLD",
			input:          "homail.aftership",
			expectedDomain: "hotmail.aftership",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := verifier.SuggestDomain(tt.input)

			assert.Equal(t, tt.expectedDomain, got)
		})
	}
}
