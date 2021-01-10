package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArgs(t *testing.T) {

	type testCase struct {
		args     []string
		expected *CommandOptions
		message  string
	}

	var testCases = []testCase{
		{
			[]string{"Title"},
			NewCommandOptions(WithOptions("Title")),
			"Extracting title from options",
		},
		{
			[]string{"Title", "+work"},
			NewCommandOptions(WithOptions("Title"), WithTags("work")),
			"Extracting title from options",
		},
		{
			[]string{"book:test"},
			NewCommandOptions(WithBook("test")),
			"Extracting title from options",
		},
	}

	for _, tc := range testCases {
		actual := ParseArgs(tc.args)
		assert.Equal(t, tc.expected, actual, tc.message)
	}

}
