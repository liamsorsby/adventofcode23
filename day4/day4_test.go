package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSumGamesPossible(t *testing.T) {
	tests := map[string]struct {
		games    string
		expected Line
	}{
		"first row": {
			games: "Card 19: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			expected: Line{
				card_number:            19,
				winning_numbers:        []int{41, 48, 83, 86, 17},
				my_numbers:             []int{83, 86, 6, 31, 17, 9, 48, 53},
				numbers_which_have_won: []int{83, 86, 17, 48},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := parseCard(test.games)

			passed := assert.Equal(t, test.expected, actual)
			if !passed {
				t.Errorf("lines possible don't match: expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestGetScore(t *testing.T) {
	tests := map[string]struct {
		lines    []Line
		expected []int
	}{
		"4 winning lines": {
			lines: []Line{
				{
					card_number:            19,
					winning_numbers:        []int{41, 48, 83, 86, 17},
					my_numbers:             []int{83, 86, 6, 31, 17, 9, 48, 53},
					numbers_which_have_won: []int{83, 86, 17, 48},
				},
			},
			expected: []int{8},
		},
		"1 winning line": {
			lines: []Line{
				{
					card_number:            19,
					winning_numbers:        []int{1, 2, 3, 4, 9},
					my_numbers:             []int{83, 86, 6, 31, 17, 9, 48, 53},
					numbers_which_have_won: []int{9},
				},
			},
			expected: []int{1},
		},
		"8 winning line": {
			lines: []Line{
				{
					card_number:            19,
					winning_numbers:        []int{1, 2, 3, 4, 5, 6, 7, 8},
					my_numbers:             []int{1, 2, 3, 4, 5, 6, 7, 8},
					numbers_which_have_won: []int{128},
				},
			},
			expected: []int{1},
		},
		"0 winning lines": {
			lines: []Line{
				{
					card_number:            19,
					winning_numbers:        []int{1, 2, 3, 4, 5, 6, 7, 8},
					my_numbers:             []int{9, 10, 11, 12, 13, 14, 15, 16},
					numbers_which_have_won: []int{},
				},
			},
			expected: []int{0},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := getLineScore(test.lines)

			passed := assert.Equal(t, test.expected, actual)
			if !passed {
				t.Errorf("lines possible don't match: expected %v, got %v", test.expected, actual)
			}
		})
	}
}
