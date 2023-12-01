package main

import (
	"bufio"
	"log"
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

func main() {
	var result int

	path, err := os.Getwd()

	file, err := os.Open(path + "/day1/resources/part1.txt")
	check(err)

	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		log.Printf("before: %s", s.Text())
		text := convertToNumbers(s.Text())
		log.Printf("before: %s", text)
		digits := getNumbers(text)
		num, err := getFirstLast(digits)
		check(err)
		result = result + num
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("result: %d", result)
}

func convertToNumbers(text string) string {
	text = strings.Replace(text, "one", "on1e", -1)
	text = strings.Replace(text, "two", "tw2o", -1)
	text = strings.Replace(text, "three", "thr3e", -1)
	text = strings.Replace(text, "four", "fo4ur", -1)
	text = strings.Replace(text, "five", "fi5ve", -1)
	text = strings.Replace(text, "six", "si6x", -1)
	text = strings.Replace(text, "seven", "sev7en", -1)
	text = strings.Replace(text, "eight", "ei8ght", -1)
	return strings.Replace(text, "nine", "ni9ne", -1)
}

func getNumbers(text string) []string {
	re := regexp.MustCompile("\\d+")
	data := re.FindAllString(text, -1)
	return strings.Split(strings.Join(data, ""), "")
}

func getFirstLast(num []string) (int, error) {
	return strconv.Atoi(num[0] + num[len(num)-1])
}
