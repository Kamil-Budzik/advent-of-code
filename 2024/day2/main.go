package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile() [][]int {
	file, err := os.Open("day2.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	defer file.Close()

	var nums [][]int
	for scanner.Scan() {
		text := scanner.Text()
		splittedLine := strings.Fields(text)

		var line []int

		for _, num := range splittedLine {
			intNum, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println(err)
				panic("Failed to convert num")
			}
			line = append(line, intNum)
		}
		nums = append(nums, line)
	}

	return nums

}

func isSafe(isIncreasing bool, current, next int) bool {
	if isIncreasing {
		return next-current <= 3 && next-current > 0
	}
	return current-next <= 3 && current-next > 0
}

func remove(slice []int, index int) []int {
	var newArr []int

	for i, num := range slice {
		if i == index {
			continue
		}
		newArr = append(newArr, num)
	}

	return newArr
}

func part1() {
	input := readFile()

	safeCount := 0

	for _, line := range input {

		isIncreasing := false

		if line[1] > line[0] {
			isIncreasing = true
		}

		isLineSafe := true

		for i := 0; i < len(line)-1; i++ {
			if !isSafe(isIncreasing, line[i], line[i+1]) {
				isLineSafe = false
				break
			}

		}

		if isLineSafe == true {
			safeCount++
		}

	}

	fmt.Println("Part 1 answer is: ", safeCount)

}

func part2() {
	input := readFile()
	safeCount := 0

	for _, line := range input {
		originalLine := make([]int, len(line))
		copy(originalLine, line)

		for j := 0; j < len(originalLine); j++ {
			newLine := remove(originalLine, j)

			isIncreasing := false
			if newLine[1] > newLine[0] {
				isIncreasing = true
			}

			isLineSafe := true
			for i := 0; i < len(newLine)-1; i++ {
				if !isSafe(isIncreasing, newLine[i], newLine[i+1]) {
					isLineSafe = false
					break
				}
			}

			if isLineSafe {
				safeCount++
				break
			}
		}
	}

	fmt.Println("Part 2 answer is:", safeCount)
}

func main() {
	part1()
	part2()
}
