package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile() ([]int, []int) {
	file, err := os.Open("day1.test.txt")
	if err != nil {
		panic(err)
	}

	var leftNums []int
	var rightNums []int

	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		splittedLine := strings.Split(text, "   ")

		left, _ := strconv.Atoi(splittedLine[0])
		right, _ := strconv.Atoi(splittedLine[1])
		leftNums = append(leftNums, left)
		rightNums = append(rightNums, right)
	}

	if len(leftNums) != len(rightNums) {
		panic("Arrays have different lenght")
	}

	return leftNums, rightNums
}

func part1() {
	leftNums, rightNums := readFile()

	sort.Ints(leftNums)
	sort.Ints(rightNums)

	if len(leftNums) != len(rightNums) {
		panic("Arrays have different length")
	}

	total := 0

	for i := 0; i < len(leftNums); i++ {
		total += int(math.Abs(float64(rightNums[i] - leftNums[i])))
	}

	fmt.Println("Total is", total)

}

func countAppearances(list []int) map[int]int {
	appearances := make(map[int]int)

	for i := 0; i < len(list); i++ {
		num := list[i]
		val, ok := appearances[num]

		if ok {
			appearances[num] = val + 1
		} else {
			appearances[num] = 1
		}
	}

	return appearances
}

func part2() {
	leftNums, rightNums := readFile()

	appearances := countAppearances(rightNums)

	total := 0
	for i := 0; i < len(leftNums); i++ {
		total += leftNums[i] * appearances[leftNums[i]]
	}

	fmt.Println(total)

}

func main() {
	part1()
	part2()
}
