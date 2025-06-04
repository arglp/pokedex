package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input 	string
		expected [] string
	}{
		{
		input:		"  hello   world ",
		expected:	[]string{"hello", "world"},
		},
		{
		input:		" hoW  aRE YOU?",
		expected:	[]string{"how", "are", "you?"},
		},
	}

	for _, c := range cases {

		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual length: %v | expected length: %v", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual: %v | expected: %v", word, expectedWord)
			}
		}
	}
}