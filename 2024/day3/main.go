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
	file, err := os.Open("day3.txt")
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

func extractInstructions(input string) []string {
	// match mul(X,Y), do(), and don't()
	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don\'t\(\)`)

	matches := regex.FindAllString(input, -1)

	return matches
}

func extractNumbersFromInstruction(v string, prefix string) (int, int) {
	v = strings.TrimPrefix(v, prefix+"(")
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

	return rightNumber, leftNumber
}

func part1() {
	input := readFile()

	mulValues := extractMulValues(input)

	total := 0

	for _, v := range mulValues {
		leftNumber, rightNumber := extractNumbersFromInstruction(v, "mul")
		total += leftNumber * rightNumber
	}

	fmt.Println("Part 1 answer is", total)

}

func part2() {
	input := readFile()

	const enableInstruction = "do()"
	const disableInstruction = "don't()"

	instructions := extractInstructions(input)
	isEnabled := true
	total := 0

	for _, ins := range instructions {

		switch ins {
		case enableInstruction:
			isEnabled = true
		case disableInstruction:
			isEnabled = false
		default:
			leftNumber, rightNumber := extractNumbersFromInstruction(ins, "mul")

			if isEnabled {
				total += leftNumber * rightNumber
			}
		}

	}

	fmt.Println("Part 2 answer is", total)

}

func main() {
	part1()
	part2()
}
