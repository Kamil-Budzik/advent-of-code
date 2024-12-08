package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Matrix [][]rune
type AntennasLocations map[rune][][]int

// 'X': [[0,2], [1,2]]

func readFile() Matrix {
	file, err := os.Open("day8.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	defer file.Close()
	var grid Matrix

	for scanner.Scan() {
		text := scanner.Text()
		line := strings.Fields(text)

		var row []rune

		for _, char := range line {
			row = append(row, []rune(char)...)
		}

		grid = append(grid, row)
	}

	return grid

}

func isInBounds(x, y, W, H int) bool {
	return x >= 0 && x < H && y >= 0 && y < W
}

func removeDuplicates(arr [][]int) [][]int {
	seen := make(map[string]bool)
	result := [][]int{}

	for _, subarray := range arr {
		key := fmt.Sprint(subarray)
		if !seen[key] {
			seen[key] = true
			result = append(result, subarray)
		}
	}

	return result
}

func part1() {
	grid := readFile()

	H := len(grid)
	W := len(grid[0])

	antennasLocations := make(AntennasLocations)

	for x := 0; x < H; x++ {
		var line []rune
		for y := 0; y < W; y++ {
			curr := grid[x][y]
			line = append(line, curr)
			if curr != '.' {
				antennasLocations[curr] = append(antennasLocations[curr], []int{x, y})
			}
		}
	}

	var antinodes [][]int
	for _, v := range antennasLocations {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				curr := v[i]
				next := v[j]

				distance := []int{next[0] - curr[0], next[1] - curr[1]}

				antinode1 := []int{next[0] + distance[0], next[1] + distance[1]}
				antinode2 := []int{curr[0] - distance[0], curr[1] - distance[1]}

				if isInBounds(antinode1[0], antinode1[1], W, H) {
					antinodes = append(antinodes, antinode1)
				}
				if isInBounds(antinode2[0], antinode2[1], W, H) {
					antinodes = append(antinodes, antinode2)
				}

			}
		}

	}

	antinodes = removeDuplicates(antinodes)
	fmt.Println("Answer to part 1", len(antinodes))

}

func main() {
	part1()
}
