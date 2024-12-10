package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Matrix [][]int

func readFile() Matrix {
	file, err := os.Open("day10.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	defer file.Close()
	var grid Matrix

	for scanner.Scan() {
		text := scanner.Text()
		line := strings.TrimSpace(text)

		var row []int

		for _, el := range line {
			num, _ := strconv.Atoi(string(el))

			row = append(row, num)
		}

		grid = append(grid, row)
	}

	return grid

}

func isInBounds(x, y, W, H int) bool {
	return x >= 0 && x < H && y >= 0 && y < W
}

var directions = [8][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

// It will return the grid position of 9, if the trailhead is correct
func createTrailHead(grid Matrix, x, y int, visited map[[2]int]bool) [][]int {
	if grid[x][y] == 9 {
		return [][]int{{x, y}}
	}

	visited[[2]int{x, y}] = true
	reachableNines := [][]int{}

	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]

		if isInBounds(newX, newY, len(grid[0]), len(grid)) &&
			!visited[[2]int{newX, newY}] &&
			grid[newX][newY]-grid[x][y] == 1 {

			result := createTrailHead(grid, newX, newY, visited)
			reachableNines = append(reachableNines, result...)
		}
	}

	return reachableNines
}

func part1() {
	grid := readFile()

	H := len(grid)
	W := len(grid[0])
	total := 0

	for x := 0; x < H; x++ {
		for y := 0; y < W; y++ {
			if grid[x][y] == 0 {
				visited := make(map[[2]int]bool)
				reachableNines := createTrailHead(grid, x, y, visited)

				uniqueNines := make(map[[2]int]bool)
				for _, pos := range reachableNines {
					uniqueNines[[2]int{pos[0], pos[1]}] = true
				}

				score := len(uniqueNines)
				total += score

			}
		}
	}

	fmt.Println("Answer to part 1", total)
}

func dfs(grid Matrix, x, y, currentHeight int, visited map[[2]int]bool) int {
	H := len(grid)
	W := len(grid[0])

	if !isInBounds(x, y, W, H) || visited[[2]int{x, y}] || grid[x][y] != currentHeight {
		return 0
	}

	if currentHeight == 9 {
		return 1
	}

	visited[[2]int{x, y}] = true
	totalPaths := 0

	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]
		totalPaths += dfs(grid, newX, newY, currentHeight+1, visited)
	}

	visited[[2]int{x, y}] = false

	return totalPaths
}

func part2() {
	grid := readFile()

	H := len(grid)
	W := len(grid[0])
	total := 0

	for x := 0; x < H; x++ {
		for y := 0; y < W; y++ {
			if grid[x][y] == 0 {
				visited := make(map[[2]int]bool)
				total += dfs(grid, x, y, 0, visited)
			}
		}
	}

	fmt.Println("Answer to part 2:", total)
}

func main() {
	part1()
	part2()
}
