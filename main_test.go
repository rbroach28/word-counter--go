package main_test

import (
	"testing"

	wordcounter "github.com/rbroach28/wordcounter"
)

func TestCountWords(t *testing.T) {

	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "five words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "empty input",
			input: "",
			wants: 0,
		},
		{
			name:  "single space",
			input: " ",
			wants: 0,
		},
		{
			name: "new lines",
			input: `one two three
			four five`,
			wants: 5,
		},
		{
			name:  "prefixed multiple spaces",
			input: "   one   two   three   ",
			wants: 3,
		},
		{
			name:  "suffixed multiple spaces",
			input: "one two three   ",
			wants: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := wordcounter.CountWords([]byte(tc.input))
			if result != tc.wants {
				t.Logf("For input '%s', expected: %d, got: %d", tc.input, tc.wants, result)
				t.Fail()
			}
		})

	}
}
