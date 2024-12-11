package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile() []int {
	file, err := os.Open("day11.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	defer file.Close()

	var stones []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		for _, el := range parts {
			stone, err := strconv.Atoi(el)
			if err != nil {
				fmt.Println("Failed to convert number", err)
			}

			stones = append(stones, stone)
		}
	}

	return stones
}

func splitStone(strStone string) (int, int) {
	mid := len(strStone) / 2

	left, err := strconv.Atoi(strStone[:mid])
	if err != nil {
		fmt.Println("failed to convert left number in SplitStone helper", err)
	}

	right, err := strconv.Atoi(strStone[mid:])
	if err != nil {
		fmt.Println("failed to convert right number in SplitStone helper", err)
	}

	return left, right
}

func part1() {
	stones := readFile()
	const blinks = 25

	blink := 0
	for blink < blinks {
		var newStones []int

		for i := 0; i < len(stones); i++ {
			stone := stones[i]
			strStone := strconv.Itoa(stone)

			if stone == 0 {
				newStones = append(newStones, 1)
			} else if len(strStone)%2 == 0 {
				left, right := splitStone(strStone)
				newStones = append(newStones, left, right)
			} else {
				newStones = append(newStones, stones[i]*2024)
			}
		}

		stones = newStones

		blink++

	}

	fmt.Println("Answer to part 1:", len(stones))

}

func main() {
	part1()
}
