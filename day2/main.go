package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	cubeRe = regexp.MustCompile(`(\d+) (blue|red|green)`)
)

func part1(question map[string]int) any {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var games []map[string]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var (
			text = scanner.Text()
			game = make(map[string]int, 3)
		)

		rounds := strings.Split(text, ": ")[1]
		for _, round := range strings.Split(rounds, ";") {
			matches := cubeRe.FindAllStringSubmatch(round, -1)
			for _, match := range matches {
				i, _ := strconv.Atoi(match[1])

				if v, ok := game[match[2]]; ok && v > i {
					continue
				}

				game[match[2]] = i
			}
		}

		games = append(games, game)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var result int
	for i, game := range games {
		if question["red"] >= game["red"] && question["green"] >= game["green"] && question["blue"] >= game["blue"] {
			result += i + 1
		}
	}

	return result
}

func part2() any {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var games []map[string]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var (
			text = scanner.Text()
			game = make(map[string]int, 3)
		)

		rounds := strings.Split(text, ": ")[1]
		for _, round := range strings.Split(rounds, ";") {
			matches := cubeRe.FindAllStringSubmatch(round, -1)
			for _, match := range matches {
				i, _ := strconv.Atoi(match[1])

				if v, ok := game[match[2]]; ok && v > i {
					continue
				}

				game[match[2]] = i
			}
		}

		games = append(games, game)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var result int
	for _, game := range games {
		total := 1
		for _, value := range game {
			total *= value
		}

		result += total
	}

	return result
}

func main() {
	fmt.Printf("Part 1: %v\n", part1(map[string]int{"red": 12, "green": 13, "blue": 14}))
	fmt.Printf("Part 2: %v\n", part2())
}
