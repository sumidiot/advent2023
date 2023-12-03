package main

import (
	"bufio"
	"day03/part1/helpers"
	"fmt"
	"os"
)

/**
 * Failed attempts:
 *   544825 - too high
 *   543348 - too high // thought negative numbers count and get subtracted
 *   538763 - too low // thought i had an overly generous right-side boundary
 *   turns out negatives weren't a thing, so it was just my bad indexing
 */

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	fmt.Println(helpers.Solve(lines))

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	file.Close()
}
