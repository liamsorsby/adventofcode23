package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Card struct {
	lines []Line
}

type Line struct {
	card_number            int
	winning_numbers        []int
	my_numbers             []int
	numbers_which_have_won []int
}

func main() {
	path, err := os.Getwd()

	file, err := os.Open(path + "/day4/resources/input.txt")
	check(err)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	card := Card{}
	lines := []Line{}

	for scanner.Scan() {
		go func(text string) {
			c := parseCard(text)
			lines = append(lines, c)
		}(scanner.Text())
	}

	card.lines = lines

	scores := getLineScore(card.lines)

	finalScore := 0
	for _, i := range scores {
		finalScore += i
	}

	fmt.Printf("final score: %d", finalScore)
}

func getLineScore(lines []Line) []int {
	scores := []int{}

	for _, i := range lines {
		go func(i Line) {
			numWinningNums := len(i.numbers_which_have_won)
			score := int(1 * math.Pow(2, float64(numWinningNums-1)))
			scores = append(scores, score)
		}(i)
	}

	return scores
}

func parseCard(row string) Line {
	line := Line{}

	space := regexp.MustCompile(`\s+`)
	singleSpace := space.ReplaceAllString(row, " ")
	initialSplit := strings.Split(singleSpace, ":")
	cardNumber, err := strconv.Atoi(splitString(initialSplit[0])[1])
	check(err)
	numbers := strings.Split(initialSplit[1], "|")
	winningNumbers := convertStringArrayToIntArray(splitString(trimSpaces(numbers[0])))
	myNumbers := convertStringArrayToIntArray(splitString(trimSpaces(numbers[1])))

	line.card_number = cardNumber
	line.my_numbers = myNumbers
	line.winning_numbers = winningNumbers
	line.numbers_which_have_won = intersection(winningNumbers, myNumbers)

	return line
}

func trimSpaces(text string) string {
	return strings.Trim(text, " ")
}

func splitString(text string) []string {
	return strings.Split(text, " ")
}

func convertStringArrayToIntArray(stringArray []string) []int {
	var intArray = []int{}
	for _, i := range stringArray {
		item, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intArray = append(intArray, item)
	}

	return intArray
}

func intersection(s1, s2 []int) (inter []int) {
	hash := make(map[int]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			inter = append(inter, e)
		}
	}
	//Remove dups from slice.
	inter = removeDups(inter)
	return
}

// Remove dups from slice.
func removeDups(elements []int) (nodups []int) {
	encountered := make(map[int]bool)
	for _, element := range elements {
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	return
}
