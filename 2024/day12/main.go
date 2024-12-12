package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Matrix [][]rune

func readFile() Matrix {
	file, err := os.Open("day12.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid Matrix

	for scanner.Scan() {
		text := scanner.Text()
		line := strings.TrimSpace(text)
		var row []rune
		for _, el := range line {
			row = append(row, el)
		}
		grid = append(grid, row)
	}

	return grid
}

func isValid(matrix Matrix, visited [][]bool, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(matrix) && y < len(matrix[0]) && !visited[x][y]
}

var directions = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func dfs(matrix Matrix, visited [][]bool, x, y int, target rune, area *int, perimeter *int) {
	visited[x][y] = true
	*area += 1

	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if nx < 0 || ny < 0 || nx >= len(matrix) || ny >= len(matrix[0]) || matrix[nx][ny] != target {
			*perimeter += 1
		} else if !visited[nx][ny] {
			dfs(matrix, visited, nx, ny, target, area, perimeter)
		}
	}
}

func calculateRegions(matrix Matrix) int {
	H, W := len(matrix), len(matrix[0])
	visited := make([][]bool, H)
	for i := range visited {
		visited[i] = make([]bool, W)
	}

	totalPrice := 0

	for x := 0; x < H; x++ {
		for y := 0; y < W; y++ {
			if !visited[x][y] {
				plant := matrix[x][y]
				area, perimeter := 0, 0
				dfs(matrix, visited, x, y, plant, &area, &perimeter)
				price := area * perimeter
				totalPrice += price
			}
		}
	}

	return totalPrice
}

func part1() {
	matrix := readFile()
	totalPrice := calculateRegions(matrix)
	fmt.Println("Answer to part 1", totalPrice)
}

func main() {
	part1()
}
