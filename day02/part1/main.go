package main

import (
	"bufio"
	"day02/part1/helpers"
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
	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		gameSample := helpers.ParseSample(line)
		fmt.Println(gameSample)
		if helpers.SupportedBy(gameSample, limits) {
			sum += gameSample.Id
		}
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	file.Close()
}
