package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile() [][]rune {
	file, err := os.Open("day4.test.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	defer file.Close()

	var input [][]rune
	for scanner.Scan() {
		text := scanner.Text()
		input = append(input, []rune(text))
	}

	return input

}

var directions = [8][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
	{-1, -1},
	{-1, 1},
	{1, -1},
	{1, 1},
}

func checkBounds(rows, cols, x, y int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func searchWord(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])

	matches := 0

	for x, row := range grid {
		for y := range row {
			for _, d := range directions {
				match := true

				for i := 0; i < len(word); i++ {
					newX := x + i*d[0]
					newY := y + i*d[1]

					if !checkBounds(rows, cols, newX, newY) || grid[newX][newY] != rune(word[i]) {
						match = false
						break
					}
				}

				if match {
					matches++
				}
			}
		}
	}

	return matches
}

func part1() {
	grid := readFile()

	matches := searchWord(grid, "XMAS")
	fmt.Println("Part 1 answer:", matches)
}

func main() {
	part1()
}
