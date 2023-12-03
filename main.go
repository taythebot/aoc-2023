package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

// Pair ...
type Pair struct {
	Value string
	Start int
	End   int
	Line  int
}

func part1() any {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var (
		line    = 0
		scanner = bufio.NewScanner(file)
		results []Pair
		symbols = make(map[int]map[int]bool)
	)
	for scanner.Scan() {
		var (
			text  = scanner.Text()
			pairs []Pair
			row   = make(map[int]bool)
			start = -1
		)

		for i, c := range text {
			var (
				isPeriod = c == 46
				isDigit  = unicode.IsDigit(c)
			)

			// Record symbol position
			if !isPeriod && !isDigit {
				row[i+1] = true
			}

			// Create pair for numbers
			if !isDigit && start != -1 {
				pairs = append(pairs, Pair{
					Value: string(text[start:i]),
					Start: start + 1,
					End:   i,
					Line:  line,
				})

				start = -1
			}

			// Start new pair
			if start == -1 && !isPeriod && isDigit {
				start = i
			}
		}

		// Add remaining pair
		if start != -1 {
			pairs = append(pairs, Pair{
				Value: string(text[start:]),
				Start: start + 1,
				End:   len(text) - 1,
				Line:  line,
			})
		}

		if len(row) > 0 {
			symbols[line] = row
		}

		results = append(results, pairs...)
		line++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var sum int
	for _, result := range results {
		var found bool

		for !found {
			// Check current row
			if row, ok := symbols[result.Line]; ok {
				if _, ok := row[result.Start-1]; ok {
					found = true
				}

				if _, ok := row[result.End+1]; ok {
					found = true
				}
			}

			// Check previous row
			if row, ok := symbols[result.Line-1]; ok {
				for i := result.Start - 1; i <= result.End+1; i++ {
					if _, ok := row[i]; ok {
						found = true
					}
				}
			}

			// Always check next row
			if row, ok := symbols[result.Line+1]; ok {
				for i := result.Start - 1; i <= result.End+1; i++ {
					if _, ok := row[i]; ok {
						found = true
					}
				}
			}

			break
		}

		if found {
			iValue, _ := strconv.Atoi(result.Value)
			sum += iValue
		}
	}

	return sum
}

func findPair(pairs []Pair, row, value int) int {
	for _, pair := range pairs {
		if pair.Line == row && (pair.Start == value || pair.End == value) {
			i, _ := strconv.Atoi(pair.Value)
			return i
		}
	}

	return 0
}

func part2() any {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var (
		line    = 0
		scanner = bufio.NewScanner(file)
		results []Pair
		symbols = make(map[int]map[int]bool)
	)
	for scanner.Scan() {
		var (
			text  = scanner.Text()
			pairs []Pair
			row   = make(map[int]bool)
			start = -1
		)

		for i, c := range text {
			var (
				isPeriod = c == 46
				isDigit  = unicode.IsDigit(c)
			)

			// Record symbol position
			if !isPeriod && !isDigit {
				row[i+1] = true
			}

			// Create pair for numbers
			if !isDigit && start != -1 {
				pairs = append(pairs, Pair{
					Value: string(text[start:i]),
					Start: start + 1,
					End:   i,
					Line:  line,
				})

				start = -1
			}

			// Start new pair
			if start == -1 && !isPeriod && isDigit {
				start = i
			}
		}

		// Add remaining pair
		if start != -1 {
			pairs = append(pairs, Pair{
				Value: string(text[start:]),
				Start: start + 1,
				End:   len(text) - 1,
				Line:  line,
			})
		}

		if len(row) > 0 {
			symbols[line] = row
		}

		results = append(results, pairs...)
		line++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var sum int
	for row, currentRow := range symbols {
		for position := range currentRow {
			pairs := make(map[int]bool)

			for len(pairs) <= 2 {
				// Check previous row
				for i := position - 1; i <= position+1; i++ {
					if pair := findPair(results, row-1, i); pair > 0 {
						pairs[pair] = true
					}
				}

				// Check current row
				for i := position - 1; i <= position+1; i++ {
					if pair := findPair(results, row, i); pair > 0 {
						pairs[pair] = true
					}
				}

				// Check next row
				for i := position - 1; i <= position+1; i++ {
					if pair := findPair(results, row+1, i); pair > 0 {
						pairs[pair] = true
					}
				}

				break
			}

			if len(pairs) == 2 {
				ratio := 1
				for pair := range pairs {
					ratio *= pair
				}

				sum += ratio
			}
		}
	}

	return sum
}

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}
