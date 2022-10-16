package emailverifier

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		inputErr       error
		expectedErrMsg string
		expectedNil    bool
	}{
		{
			name:           "Parse 550RCPTError",
			inputErr:       errors.New("550 This mailbox does not exist"),
			expectedErrMsg: ErrServerUnavailable,
		},
		{
			name:           "Parse 550BlockedRCPTError",
			inputErr:       errors.New("550 spamhaus"),
			expectedErrMsg: ErrBlocked,
		},
		{
			name:           "Parse ConnectMailExchangerError",
			inputErr:       errors.New("Timeout connecting to mail-exchanger"),
			expectedErrMsg: ErrTimeout,
		},
		{
			name:           "Parse NoMxRecordsFoundError",
			inputErr:       errors.New("No MX records found"),
			expectedErrMsg: "No MX records found",
		},
		{
			name:           "Parse FullInBoxError",
			inputErr:       errors.New("452 full Inbox"),
			expectedErrMsg: ErrFullInbox,
		},
		{
			name:           "Parse DailSMTPServerError",
			inputErr:       errors.New("Unexpected response dialing SMTP server"),
			expectedErrMsg: "Unexpected response dialing SMTP server",
		},
		{
			name:           "Parse Error_Code550",
			inputErr:       errors.New("550"),
			expectedErrMsg: ErrServerUnavailable,
		},
		{
			name:           "Parse Error_Code400_Nil",
			inputErr:       errors.New("400"),
			expectedErrMsg: "",
			expectedNil:    true,
		},
		{
			name:           "Parse ParseError_Code401",
			inputErr:       errors.New("401"),
			expectedErrMsg: "401",
		},
		{
			name:           "Parse Error_Code421",
			inputErr:       errors.New("421"),
			expectedErrMsg: ErrTryAgainLater,
		},
		{
			name:           "Parse Error_Code450",
			inputErr:       errors.New("450"),
			expectedErrMsg: ErrMailboxBusy,
		},
		{
			name:           "Parse Error_Code451",
			inputErr:       errors.New("451"),
			expectedErrMsg: ErrExceededMessagingLimits,
		},
		{
			name:           "Parse Error_Code452",
			inputErr:       errors.New("452"),
			expectedErrMsg: ErrTooManyRCPT,
		},
		{
			name:           "Parse Error_Code503",
			inputErr:       errors.New("503"),
			expectedErrMsg: ErrNeedMAILBeforeRCPT,
		},
		{
			name:           "Parse Error_Code551",
			inputErr:       errors.New("551"),
			expectedErrMsg: ErrRCPTHasMoved,
		},
		{
			name:           "Parse Error_Code552",
			inputErr:       errors.New("552"),
			expectedErrMsg: ErrFullInbox,
		},
		{
			name:           "Parse Error_Code553",
			inputErr:       errors.New("553"),
			expectedErrMsg: ErrNoRelay,
		},
		{
			name:           "Parse Error_Code554",
			inputErr:       errors.New("554"),
			expectedErrMsg: ErrNotAllowed,
		},
		{
			name:           "Parse Error_basicErr_timeout",
			inputErr:       errors.New("559 timeout"),
			expectedErrMsg: ErrTimeout,
		},
		{
			name:           "Parse Error_basicErr_blocked",
			inputErr:       errors.New("559 blocked"),
			expectedErrMsg: ErrBlocked,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ParseSMTPError(tt.inputErr)
			if !tt.expectedNil {
				assert.Equal(t, tt.expectedErrMsg, got.Message)
			}
		})
	}
}
