package main

import (
	"bufio"
	helpers1 "day02/part1/helpers"
	helpers2 "day02/part2/helpers"
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
		gameSample := helpers1.ParseSample(line)
		support := helpers2.FindSupport(gameSample)
		power := helpers2.Power(support)
		sum += power
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	file.Close()
}
