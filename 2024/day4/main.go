package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile() [][]rune {
	file, err := os.Open("day4.txt")
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

func checkXmasPattern(topLeft, topRight, bottomLeft, bottomRight rune) bool {
	if (topLeft == 'M' && topRight == 'M' && bottomLeft == 'S' && bottomRight == 'S') ||
		(topLeft == 'S' && topRight == 'S' && bottomLeft == 'M' && bottomRight == 'M') ||
		(topLeft == 'M' && topRight == 'S' && bottomLeft == 'M' && bottomRight == 'S') ||
		(topLeft == 'S' && topRight == 'M' && bottomLeft == 'S' && bottomRight == 'M') {
		return true
	}
	return false

}

func searchXmas(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])

	matches := 0

	for x, row := range grid {
		for y := range row {
			if grid[x][y] != 'A' {
				continue
			}

			topLeftX, topLeftY := x-1, y-1
			topRightX, topRightY := x-1, y+1
			bottomLeftX, bottomLeftY := x+1, y-1
			bottomRightX, bottomRightY := x+1, y+1

			if !checkBounds(rows, cols, topLeftX, topLeftY) ||
				!checkBounds(rows, cols, topRightX, topRightY) ||
				!checkBounds(rows, cols, bottomLeftX, bottomLeftY) ||
				!checkBounds(rows, cols, bottomRightX, bottomRightY) {
				continue
			}

			topLeft := grid[topLeftX][topLeftY]
			topRight := grid[topRightX][topRightY]
			bottomLeft := grid[bottomLeftX][bottomLeftY]
			bottomRight := grid[bottomRightX][bottomRightY]

			if checkXmasPattern(topLeft, topRight, bottomLeft, bottomRight) {
				matches++
			}
		}
	}

	return matches
}

func part2() {
	grid := readFile()

	matches := searchXmas(grid)
	fmt.Println("Part 2 answer:", matches)
}

func main() {
	part1()
	part2()
}
