package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Set struct {
	green int
	blue  int
	red   int
}

type Game struct {
	index int
	set   []Set
}

type Games struct {
	Games []Game
}

func main() {
	actualSet := Set{green: 13, blue: 14, red: 12}
	path, err := os.Getwd()

	file, err := os.Open(path + "/day2/resources/input.txt")
	check(err)

	defer file.Close()
	s := bufio.NewScanner(file)
	game := []Game{}

	for s.Scan() {
		g := convertTextToStruct(s.Text())
		game = append(game, g)
	}

	games := Games{Games: game}
	lowestPossible := getSmallestPossibleInEachBag(games)
	power := getPowerFromSets(lowestPossible)

	possibe := checkGamesPossible(games, actualSet)
	sumOf := getSumGamesPossible(possibe)
	fmt.Printf("Sum of indexes of possible games: %d. Power of lowest possible: %d", sumOf, power)
}

func getSumGamesPossible(games Games) int {
	sumOfIndexes := 0
	for _, v := range games.Games {
		sumOfIndexes += v.index
	}
	return sumOfIndexes
}

func getPowerFromSets(sets []Set) int {
	result := 0
	for _, set := range sets {
		value := set.red * set.green * set.blue
		result += value
	}

	return result
}

func getSmallestPossibleInEachBag(games Games) []Set {
	sets := []Set{}
	for _, game := range games.Games {
		smallestGreen := 0
		smallestRed := 0
		smallestBlue := 0

		for _, set := range game.set {
			if set.green > smallestGreen {
				smallestGreen = set.green
			}
			if set.red > smallestRed {
				smallestRed = set.red
			}
			if set.blue > smallestBlue {
				smallestBlue = set.blue
			}
		}
		sets = append(sets, Set{red: smallestRed, blue: smallestBlue, green: smallestGreen})
	}

	return sets
}

func checkGamesPossible(games Games, possible Set) Games {
	possibleGames := []Game{}

	for _, game := range games.Games {
		isPossible := true
		for _, set := range game.set {
			if set.green > possible.green {
				isPossible = false
				break
			}
			if set.red > possible.red {
				isPossible = false
				break
			}
			if set.blue > possible.blue {
				isPossible = false
				break
			}
		}
		if isPossible {
			possibleGames = append(possibleGames, game)
		}
	}

	return Games{Games: possibleGames}
}

func convertTextToStruct(text string) Game {
	game := Game{}

	data := strings.Split(text, ":")

	// parse index of the game
	gameData := strings.Split(data[0], " ")
	index, err := strconv.Atoi(gameData[1])
	check(err)
	game.index = index

	// parse sets
	setsText := strings.Trim(data[1], " ")
	setsChunks := strings.Split(setsText, ";")
	sets := []Set{}

	for _, v := range setsChunks {
		setsData := strings.Split(v, ",")
		set := Set{}

		for _, i := range setsData {
			s := strings.Trim(i, " ")
			item := strings.Split(s, " ")
			value, err := strconv.Atoi(item[0])
			check(err)

			switch item[1] {
			case "blue":
				set.blue = value
				break
			case "green":
				set.green = value
				break
			case "red":
				set.red = value
				break
			}
		}
		sets = append(sets, set)

	}
	game.set = sets

	return game
}
