package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input 		string
		expected 	[]string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("the length of the result of cleanInput(%v) doesn't match the length of the expected slice(%v)", len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("words don't match. expected: %s - got: %s ", expectedWord, word)
			}
		}
	}
}
