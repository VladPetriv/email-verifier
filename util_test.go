package emailverifier

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDomain(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedDomain string
	}{
		{
			name:           "Ok",
			input:          "yahoo.com.is",
			expectedDomain: "com.is",
		},
		{
			name:           "With upper case",
			input:          "YaHoO.cOm",
			expectedDomain: "yahoo.com",
		},
		{
			name:           "Make sense",
			input:          "t.example.yahoo.com",
			expectedDomain: "yahoo.com",
		},
		{
			name:           "Empty string",
			input:          "",
			expectedDomain: "",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := parsedDomain(tt.input)
			assert.Equal(t, tt.expectedDomain, got)
		})
	}
}

func TestDomainToASCII(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedResult string
	}{
		{
			name:           "Ok",
			input:          "testingΣ✪✯☭➳卐.org",
			expectedResult: "xn--testing-0if2960fjccubz8h9z13a.org",
		},
		{
			name:           "With normal domain",
			input:          "testing.org",
			expectedResult: "testing.org",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := domainToASCII(tt.input)
			assert.Equal(t, tt.expectedResult, got)
		})
	}
}

func TestSplitDomain(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedResult [2]string
	}{
		{
			name:           "Ok",
			input:          "aftership.com",
			expectedResult: [2]string{"aftership", "com"},
		},
		{
			name:           "domain with NoSLD",
			input:          "com",
			expectedResult: [2]string{"", "com"},
		},
		{
			name:           "domain with nil string",
			input:          "",
			expectedResult: [2]string{"", ""},
		},
		{
			name:           "domain with sub domain",
			input:          "develop.aftership.com",
			expectedResult: [2]string{"aftership", "com"},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sld, tld := splitDomain(tt.input)
			assert.Equal(t, tt.expectedResult[0], sld)
			assert.Equal(t, tt.expectedResult[1], tld)
		})
	}
}

func TestCallJobFuncWithParams_NoOutput(t *testing.T) {
	f := func(a string) { fmt.Println(a) }
	ret := callJobFuncWithParams(f, []interface{}{"testing"})
	assert.Nil(t, ret)
}

func TestCallJobFuncWithParams_WithOutput(t *testing.T) {
	f := func(a int) int { return a * (1 + a) }
	ret := callJobFuncWithParams(f, []interface{}{2})
	assert.Equal(t, int64(6), ret[0].Int())
}

func TestCallJobFuncForgetParams(t *testing.T) {
	f := func(a int) int { return a * (1 + a) }
	ret := callJobFuncWithParams(f, nil)
	assert.Nil(t, ret)
}

func TestCallJobFuncWithWrongFunc(t *testing.T) {
	f := 3
	ret := callJobFuncWithParams(f, nil)
	assert.Nil(t, ret)
}
