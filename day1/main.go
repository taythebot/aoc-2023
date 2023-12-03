package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func part1() any {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var results int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var (
			text    = []rune(scanner.Text())
			numbers string
		)

		for _, c := range text {
			if unicode.IsDigit(c) {
				numbers = string(c)
				break
			}
		}

		for i := len(text) - 1; i >= 0; i-- {
			if unicode.IsDigit(text[i]) {
				numbers += string(text[i])
				break
			}
		}

		i, _ := strconv.Atoi(numbers)
		results += i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return results
}

var (
	numbers = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
)

func checkNumber(text string) string {
	for token, value := range numbers {
		if strings.Contains(text, token) {
			return value
		}
	}

	return ""
}

func part2() any {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var results int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var (
			text    = []rune(scanner.Text())
			numbers string
			buffer  string
		)

		for _, c := range text {
			if unicode.IsDigit(c) {
				numbers = string(c)
				break
			}

			buffer += string(c)

			if len(buffer) >= 3 {
				if value := checkNumber(buffer); value != "" {
					numbers += value
					break
				}
			}
		}

		buffer = ""

		for i := len(text) - 1; i >= 0; i-- {
			if unicode.IsDigit(text[i]) {
				numbers += string(text[i])
				break
			}

			// Reverse add to buffer
			buffer = string(text[i]) + buffer

			if len(buffer) >= 3 {
				if value := checkNumber(buffer); value != "" {
					numbers += value
					break
				}
			}
		}

		i, _ := strconv.Atoi(numbers)
		results += i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return results
}

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}
