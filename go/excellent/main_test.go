package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	testCases := []struct {
		number   int
		expected string
	}{
		{1, "Odd"},
		{2, "Even"},
		{3, "Odd"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			actual := EvenOrOdd(tc.number)
			if actual != tc.expected {
				t.Errorf("EvenOrOdd(%d) = %s; expected %s", tc.number, actual, tc.expected)
			}
		})
	}
}
