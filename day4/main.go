package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseNumbers(numbers string) map[int]bool {
	results := make(map[int]bool)
	for _, n := range strings.Split(numbers, " ") {
		if n == "" {
			continue
		}

		i, _ := strconv.Atoi(n)
		results[i] = true
	}

	return results
}

func part1() any {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var (
		scanner = bufio.NewScanner(file)
		sum     int
	)
	for scanner.Scan() {
		var (
			text  = scanner.Text()
			score int
		)

		split := strings.Split(text, "|")

		numbers := parseNumbers(strings.Split(split[0], ":")[1])
		winningNumbers := parseNumbers(split[1])

		for n := range numbers {
			if _, ok := winningNumbers[n]; ok {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}

		sum += score
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

type Card struct {
	Numbers        map[int]bool
	WinningNumbers map[int]bool
	Copies         int
}

func part2() any {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var (
		scanner = bufio.NewScanner(file)
		cards   []Card
	)
	for i := 1; scanner.Scan(); i++ {
		split := strings.Split(scanner.Text(), "|")

		numbers := parseNumbers(strings.Split(split[0], ":")[1])
		winningNumbers := parseNumbers(split[1])

		cards = append(cards, Card{
			Numbers:        numbers,
			WinningNumbers: winningNumbers,
			Copies:         1,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var totalCards int
	for i, card := range cards {
		var matching int
		for n := range card.Numbers {
			if _, ok := card.WinningNumbers[n]; ok {
				matching++
			}
		}

		if matching > 0 {
			for x := i + 1; x <= i+matching && x < len(cards); x++ {
				cards[x].Copies += card.Copies
			}
		}

		totalCards += card.Copies
	}

	return totalCards
}

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}
