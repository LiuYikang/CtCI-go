package stringandarray

import (
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"abcd", "dcba"},
		{"abcc", "ccba"},
		{" ", " "},
		{"", ""},
	}
	for _, c := range cases {
		actual := Reverse(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s. Expected: %s, actual: %s\n", c.input, c.expected, actual)
		}
	}
}
