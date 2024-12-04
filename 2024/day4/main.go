package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
	{-1, -1},
	{-1, 1},
	{1, -1},
	{1, 1},
}

func readFile() []string {
	file, err := os.Open("day4.test.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	defer file.Close()

	var input []string
	for scanner.Scan() {
		text := scanner.Text()
		input = append(input, text)
	}

	return input

}

func checkBounds(rows, cols, x, y int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func part1() {
	grid := readFile()
	for _, row := range grid {
		fmt.Println(row)
	}

	rows := len(grid)
	cols := len(grid[0])

	for x, row := range grid {
		for y, col := range row {
			if grid[x][y] != byte('K') {
				continue
			}
			_ = col
			for _, d := range directions {
				newX := x + d[0]
				newY := y + d[1]

				if checkBounds(rows, cols, newX, newY) {
					fmt.Printf("Value at (%d, %d): %s\n", newX, newY, string(grid[newX][newY]))
				}
			}
		}
	}

}

func main() {
	part1()
}
