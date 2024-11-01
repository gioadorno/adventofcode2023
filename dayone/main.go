package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var wordToNumber = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"zero":  "0",
}

func main() {
	sum := 0

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		sum += findNumbers(line)
	}

	fmt.Printf("Sum: %d", sum)

}

func findNumber(line string, reverse bool) string {
	var result string

	var start, end, step int
	if reverse {
		start = len(line) - 1
		end = -1
		step = -1
	} else {
		start = 0
		end = len(line)
		step = 1
	}

	for i := start; i != end; i += step {
		character := line[i]
		for word, number := range wordToNumber {
			if reverse {
				if strings.HasSuffix(line[:i+1], word) {
					result = number
					break
				}
			} else if strings.HasPrefix(line[i:], word) {
				result = number
				break
			}
		}
		if result == "" && character >= '0' && character <= '9' {
			result = string(character)
		}
		if result != "" {
			break
		}
	}

	return result
}

func findNumbers(line string) int {
	first := findNumber(line, false)
	last := findNumber(line, true)

	fmt.Println(line)
	fmt.Println(first, last)

	n, _ := strconv.Atoi(first + last)

	return n
}
