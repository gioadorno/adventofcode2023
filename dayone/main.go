package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func findNumbers(line string) int {
	lineNumber := ""
	firstNumber := ""
	lastNumber := ""

	for i := 0; i < len(line); i++ {
		str := string(line[i])
		_, err := strconv.Atoi(str)
		if err == nil {
			if firstNumber == "" {
				firstNumber = str
			}
			lastNumber = str
		}
	}
	if lastNumber == "" {
		lastNumber = firstNumber
	}
	lineNumber = firstNumber + lastNumber
	fmt.Println(lineNumber)
	n, _ := strconv.Atoi(lineNumber)
	return n
}
