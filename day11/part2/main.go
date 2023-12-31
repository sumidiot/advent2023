package main

import (
	"bufio"
	"day11/part2/helpers"
	"fmt"
	"os"
)

/**
 * Failed attempts:
 *   * 82000210 too low
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

	fmt.Println(helpers.Solve(lines, 1000000))

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	file.Close()
}
