package emailverifier

import (
	"testing"
)

func TestCheckAddressSyntax(t *testing.T) {
	t.Parallel()

	tests := []struct {
		email  string
		format bool
	}{
		{email: "example@domain.com", format: true},
		{email: "support@yahoo.com", format: true},
		{email: " jerry@gmail.com", format: false},
		{email: "tool@163.com", format: true},
		{email: "😀@gmail.com", format: false},
		{email: "user@gma3il.com", format: true},
		{email: "a_b@github.com", format: true},
		{email: "abc@доменное.com", format: true},
	}

	for _, tt := range tests {
		tt := tt

		t.Run("Validate email format", func(t *testing.T) {
			t.Parallel()

			address := verifier.ParseAddress(tt.email)
			if !address.Valid && tt.format == true {
				t.Errorf(`"%s" check failed with an unexpected error`, tt.email)
			}

			if address.Valid && tt.format == false {
				t.Errorf(`"%s" => incorrect email address`, tt.email)
			}
		})
	}
}
