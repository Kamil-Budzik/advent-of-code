package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile() string {
	file, err := os.Open("day3.test.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	defer file.Close()

	var input string
	for scanner.Scan() {
		text := scanner.Text()

		input += text
	}

	return input
}

func extractMulValues(input string) []string {
	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := regex.FindAllString(input, -1)

	return matches
}

func part1() {
	input := readFile()

	mulValues := extractMulValues(input)

	total := 0

	for _, v := range mulValues {
		v = strings.TrimPrefix(v, "mul(")
		v = strings.TrimSuffix(v, ")")
		nums := strings.Split(v, ",")

		if len(nums) != 2 {
			fmt.Println("Numbers length should be 2")
		}

		leftNumber, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println("Failed to convert Left number", err)
		}

		rightNumber, _ := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println("Failed to convert right number", err)
		}

		total += leftNumber * rightNumber
	}

	fmt.Println("Part 1 answer is", total)

}

func main() {
	part1()
}
