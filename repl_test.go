package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct{
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "the cat says meow",
			expected: []string{"the", "cat", "says", "meow"},
		},
		{
			input: "meow cat please meow back",
			expected: []string{"meow", "cat", "please", "meow", "back"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check length of actual slice against expected slice and t.Errorf if failed test
		if len(actual) != len(c.expected) {
			t.Errorf("Doesn't match")
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in slice and if they don't match, t.Errorf and fail test
			if word != expectedWord {
				t.Errorf("Word doesn't match")
				continue
			}
		}
		
	}

}