package main

import "testing"

func TestFibDP(t *testing.T) {
	cases := []struct {
		input    int
		expected int
	}{
		{
			1,
			0,
		},

		{
			10,
			55,
		},

		{
			50,
			12586269025,
		},
	}

	for _, c := range cases {
		actual := fibDP(c.input)
		if actual != c.expected {
			t.Errorf("expected %d to equal %d", actual, c.expected)
		}
	}
}
