package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const fileDir string = "day2/input"

func main() {
	fmt.Printf("AoC 2021 Part 1 - What do you get if you multiply your final horizontal position by your final depth? %v\n", partOne())
	fmt.Printf("AoC 2021 Part 2 - What do you get if you multiply your final horizontal position by your final depth? %v\n", partTwo())
}

func partOne() int64 {
	absPath, _ := os.Getwd()
	file, err := os.Open(strings.Join([]string{absPath, fileDir}, "/"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '

	var horizontal int64
	var depth int64

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		units, _ := strconv.Atoi(row[1])

		switch row[0] {
		case "forward":
			horizontal += int64(units)
		case "down":
			depth += int64(units)
		case "up":
			depth -= int64(units)
		}
	}

	return horizontal * depth
}

func partTwo() int64 {
	absPath, _ := os.Getwd()
	file, err := os.Open(strings.Join([]string{absPath, fileDir}, "/"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '

	var horizontal int64
	var depth int64
	var aim int64

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		units, _ := strconv.Atoi(row[1])

		switch row[0] {
		case "forward":
			horizontal += int64(units)
			depth += aim * int64(units)
		case "down":
			aim += int64(units)
		case "up":
			aim -= int64(units)
		}
	}

	return horizontal * depth
}
