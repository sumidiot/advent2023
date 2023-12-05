package main

import (
	"bufio"
	"day03/part2/helpers"
	"fmt"
	"os"
)

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
