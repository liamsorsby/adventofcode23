package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertTextToStruct(t *testing.T) {
	expected := Game{
		index: 1,
		set: []Set{
			{blue: 12},
			{green: 2, blue: 13, red: 19},
			{green: 3, blue: 14, red: 13},
		},
	}
	text := "Game 1: 12 blue; 2 green, 13 blue, 19 red; 13 red, 3 green, 14 blue"
	got := convertTextToStruct(text)
	assert.Equal(t, expected, got)
}

func TestGetSumGamesPossible(t *testing.T) {
	tests := map[string]struct {
		games    Games
		expected int
	}{
		"3 games": {
			games: Games{
				Games: []Game{
					{index: 1, set: []Set{
						{green: 1, blue: 8, red: 10},
						{green: 10, blue: 5, red: 1},
						{green: 5, blue: 3, red: 1},
					}},
					{index: 2, set: []Set{
						{green: 15, blue: 6, red: 4},
						{green: 2, blue: 3, red: 8},
						{green: 6, blue: 2, red: 3},
					}},
					{index: 3, set: []Set{
						{green: 55, blue: 6, red: 4},
					}},
				},
			},
			expected: 6,
		},
		"2 Games": {
			games: Games{
				Games: []Game{
					{index: 2, set: []Set{
						{green: 15, blue: 6, red: 4},
						{green: 2, blue: 13, red: 8},
						{green: 6, blue: 2, red: 3},
					}},
					{index: 3, set: []Set{
						{green: 55, blue: 100, red: 4},
					}},
				},
			},
			expected: 5,
		},
		"None": {
			games: Games{
				Games: []Game{},
			},
			expected: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := getSumGamesPossible(test.games)

			passed := assert.Equal(t, test.expected, actual)
			if !passed {
				t.Errorf("games possible don't match: expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestSmallestSets(t *testing.T) {
	tests := map[string]struct {
		games    Games
		expected []Set
	}{
		"3 games": {
			games: Games{
				Games: []Game{
					{index: 1, set: []Set{
						{green: 1, blue: 8, red: 10},
						{green: 10, blue: 5, red: 1},
						{green: 5, blue: 3, red: 1},
					}},
					{index: 2, set: []Set{
						{green: 15, blue: 6, red: 4},
						{green: 2, blue: 3, red: 8},
						{green: 6, blue: 2, red: 3},
					}},
					{index: 3, set: []Set{
						{green: 55, blue: 6, red: 4},
					}},
				},
			},
			expected: []Set{
				{green: 10, blue: 8, red: 10},
				{green: 15, blue: 6, red: 8},
				{green: 55, blue: 6, red: 4},
			},
		},
		"2 Games": {
			games: Games{
				Games: []Game{
					{index: 2, set: []Set{
						{green: 15, blue: 6, red: 4},
						{green: 2, blue: 13, red: 8},
						{green: 6, blue: 2, red: 3},
					}},
					{index: 3, set: []Set{
						{green: 55, blue: 100, red: 4},
					}},
				},
			},
			expected: []Set{
				{green: 15, blue: 13, red: 8},
				{green: 55, blue: 100, red: 4},
			},
		},
		"None": {
			games: Games{
				Games: []Game{},
			},
			expected: []Set{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := getSmallestPossibleInEachBag(test.games)

			passed := assert.Equal(t, test.expected, actual)
			if !passed {
				t.Errorf("games possible don't match: expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestPowerFromSets(t *testing.T) {
	tests := map[string]struct {
		sets     []Set
		expected int
	}{
		"3 games": {
			sets: []Set{
				{green: 10, blue: 8, red: 10}, //800
				{green: 15, blue: 6, red: 8},  //720
				{green: 55, blue: 6, red: 4},  //1320
			},
			expected: 2840,
		},
		"2 Games": {
			sets: []Set{
				{green: 15, blue: 13, red: 8},  //1560
				{green: 55, blue: 100, red: 4}, //22000
			},
			expected: 23560,
		},
		"None": {
			sets:     []Set{},
			expected: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := getPowerFromSets(test.sets)

			passed := assert.Equal(t, test.expected, actual)
			if !passed {
				t.Errorf("games possible don't match: expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestGamesArePossible(t *testing.T) {
	tests := map[string]struct {
		games    Games
		inBag    Set
		expected Games
	}{
		"Greens": {
			games: Games{
				Games: []Game{
					{index: 1, set: []Set{
						{green: 1, blue: 8, red: 10},
						{green: 10, blue: 5, red: 1},
						{green: 5, blue: 3, red: 1},
					}},
					{index: 2, set: []Set{
						{green: 15, blue: 6, red: 4},
						{green: 2, blue: 3, red: 8},
						{green: 6, blue: 2, red: 3},
					}},
					{index: 3, set: []Set{
						{green: 55, blue: 6, red: 4},
					}},
				},
			},
			inBag: Set{red: 40, blue: 50, green: 30},
			expected: Games{
				Games: []Game{
					{index: 1, set: []Set{
						{green: 1, blue: 8, red: 10},
						{green: 10, blue: 5, red: 1},
						{green: 5, blue: 3, red: 1},
					}},
					{index: 2, set: []Set{
						{green: 15, blue: 6, red: 4},
						{green: 2, blue: 3, red: 8},
						{green: 6, blue: 2, red: 3},
					}},
				},
			},
		},
		"Blues": {
			games: Games{
				Games: []Game{
					{index: 1, set: []Set{
						{green: 1, blue: 8, red: 10},
						{green: 10, blue: 5, red: 1},
						{green: 5, blue: 3, red: 1},
					}},
					{index: 2, set: []Set{
						{green: 15, blue: 6, red: 4},
						{green: 2, blue: 13, red: 8},
						{green: 6, blue: 2, red: 3},
					}},
					{index: 3, set: []Set{
						{green: 55, blue: 100, red: 4},
					}},
				},
			},
			inBag: Set{red: 56, blue: 10, green: 30},
			expected: Games{
				Games: []Game{
					{index: 1, set: []Set{
						{green: 1, blue: 8, red: 10},
						{green: 10, blue: 5, red: 1},
						{green: 5, blue: 3, red: 1},
					}},
				},
			},
		},
		"Reds": {
			games: Games{
				Games: []Game{
					{index: 1, set: []Set{
						{green: 1, blue: 8, red: 10},
						{green: 10, blue: 5, red: 1},
						{green: 5, blue: 3, red: 1},
					}},
					{index: 2, set: []Set{
						{green: 15, blue: 6, red: 1},
						{green: 2, blue: 13, red: 1},
						{green: 6, blue: 2, red: 1},
					}},
					{index: 3, set: []Set{
						{green: 55, blue: 100, red: 44},
					}},
				},
			},
			inBag: Set{red: 2, blue: 20, green: 30},
			expected: Games{
				Games: []Game{
					{index: 2, set: []Set{
						{green: 15, blue: 6, red: 1},
						{green: 2, blue: 13, red: 1},
						{green: 6, blue: 2, red: 1},
					}},
				},
			},
		},
		"None": {
			games: Games{
				Games: []Game{
					{index: 1, set: []Set{
						{green: 1, blue: 8, red: 10},
						{green: 10, blue: 5, red: 1},
						{green: 5, blue: 3, red: 1},
					}},
					{index: 2, set: []Set{
						{green: 15, blue: 6, red: 4},
						{green: 2, blue: 3, red: 8},
						{green: 6, blue: 2, red: 3},
					}},
					{index: 3, set: []Set{
						{green: 55, blue: 6, red: 4},
					}},
				},
			},
			inBag: Set{red: 0, blue: 0, green: 0},
			expected: Games{
				Games: []Game{},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := checkGamesPossible(test.games, test.inBag)

			passed := assert.Equal(t, test.expected, actual)
			if !passed {
				t.Errorf("games possible don't match: expected %v, got %v", test.expected, actual)
			}
		})
	}
}
