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

	var input strings.Builder
	for scanner.Scan() {
		input.WriteString(scanner.Text())
	}

	return input.String()
}

func extractMulValues(input string) []string {
	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := regex.FindAllString(input, -1)

	return matches
}

func extractInstructions(input string) []string {
	// match mul(X,Y), do(), and don't(), where X and Y are 1-3 digit numbers
	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don\'t\(\)`)
	matches := regex.FindAllString(input, -1)

	return matches
}

func extractNumbersFromInstruction(v string, prefix string) (int, int, error) {
	v = strings.TrimPrefix(v, prefix+"(")
	v = strings.TrimSuffix(v, ")")
	nums := strings.Split(v, ",")

	if len(nums) != 2 {
		return 0, 0, fmt.Errorf("invalid instruction format: %s", v)
	}

	leftNumber, err := strconv.Atoi(nums[0])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert left number: %s", nums[0])
	}

	rightNumber, err := strconv.Atoi(nums[1])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert right number: %s", nums[1])
	}

	return leftNumber, rightNumber, nil
}

func part1() {
	input := readFile()

	mulValues := extractMulValues(input)

	total := 0

	for _, v := range mulValues {
		leftNumber, rightNumber, err := extractNumbersFromInstruction(v, "mul")
		if err != nil {
			fmt.Println("Error extracting numbers", err)
			continue
		}
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
			leftNumber, rightNumber, err := extractNumbersFromInstruction(ins, "mul")

			if err != nil {
				fmt.Println("Skipping invalid instruction", err)
				continue
			}

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
