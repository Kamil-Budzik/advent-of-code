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
		text = strings.TrimSpace(text)
		splittedLine := strings.Split(text, " ")

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

func part1() {
	nums := readFile()

	safeCount := 0

	for _, line := range nums {

		isIncreasing := false

		// 1 element line is always safe
		if len(line) == 1 {
			safeCount += 1
			continue
		}

		if line[1] > line[0] {
			isIncreasing = true
		}

		isLineSafe := true

		for i := 0; i < len(line)-1; i++ {
			current := line[i]
			next := line[i+1]

			if isIncreasing == true {

				if !(next-current <= 3 && next-current > 0) {
					isLineSafe = false
				}

			} else {
				if !(current-next <= 3 && current-next > 0) {
					isLineSafe = false
				}

			}

		}

		if isLineSafe == true {
			safeCount++
		}

	}

	fmt.Println("Part 1 answer is: ", safeCount)

}

func main() {
	part1()
}
