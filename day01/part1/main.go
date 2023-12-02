package main

import (
	"bufio"
	"day01/part1/helpers"
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

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += helpers.Calval(line)
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	file.Close()
}
