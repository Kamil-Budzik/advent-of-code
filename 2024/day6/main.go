package main

import (
	"bufio"
	"fmt"
	"os"
)

type Matrix [][]rune

const TOP = '^'
const RIGHT = '>'
const BOTTOM = 'v'
const LEFT = '<'

func readFile() (Matrix, int, int, rune) {
	file, err := os.Open("day6.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	defer file.Close()

	var matrix Matrix
	var startX, startY int
	var startingDirection rune

	for x := 0; scanner.Scan(); x++ {
		var line []rune

		for y, char := range scanner.Text() {
			line = append(line, char)

			if char == TOP || char == RIGHT || char == BOTTOM || char == LEFT {
				startX, startY, startingDirection = x, y, char
			}
		}

		matrix = append(matrix, line)
	}

	return matrix, startX, startY, startingDirection

}

func didEscape(x, y, rows, cols int) bool {
	return x < 0 || x >= rows || y < 0 || y >= cols
}

func getNewDirection(currentDirection rune) rune {
	switch currentDirection {
	case TOP:
		return RIGHT
	case RIGHT:
		return BOTTOM
	case BOTTOM:
		return LEFT
	default: // LEFT case
		return TOP
	}
}

func didEncounterObstacle(nextPosition rune) bool {
	return nextPosition == '#'
}

func generateGuardPath(matrix Matrix, x int, y int, direction rune) Matrix {
	var newDirection rune = direction
	newX, newY := x, y

	switch direction {
	case TOP:
		newX = x - 1
	case RIGHT:
		newY = y + 1
	case BOTTOM:
		newX = x + 1
	case LEFT:
		newY = y - 1
	}

	if didEscape(newX, newY, len(matrix), len(matrix[0])) {
		matrix[x][y] = 'X'
		return matrix
	}

	if didEncounterObstacle(matrix[newX][newY]) {
		newDirection = getNewDirection(direction)
		matrix[x][y] = 'X'
		return generateGuardPath(matrix, x, y, newDirection)
	}

	matrix[newX][newY] = direction
	matrix[x][y] = 'X'
	return generateGuardPath(matrix, newX, newY, newDirection)
}

func countDistincsPositions(matrix Matrix) int {
	total := 0
	for _, rows := range matrix {
		for _, y := range rows {
			if y == 'X' {
				total++
			}
		}
	}

	return total
}

func part1() {
	matrix, startX, startY, direction := readFile()

	guardPath := generateGuardPath(matrix, startX, startY, direction)

	fmt.Println("Answer to part 1:", countDistincsPositions(guardPath))

}

func simulateGuardPath(matrix Matrix, x int, y int, direction rune, visited map[string]bool) (Matrix, bool) {
	var newDirection rune = direction
	newX, newY := x, y

	positionKey := fmt.Sprintf("%d,%d,%c", x, y, direction)

	if visited[positionKey] {
		return matrix, true
	} else {
		visited[positionKey] = true
	}

	switch direction {
	case TOP:
		newX = x - 1
	case RIGHT:
		newY = y + 1
	case BOTTOM:
		newX = x + 1
	case LEFT:
		newY = y - 1
	}

	if didEscape(newX, newY, len(matrix), len(matrix[0])) {
		matrix[x][y] = 'X'
		return matrix, false
	}

	if didEncounterObstacle(matrix[newX][newY]) {
		newDirection = getNewDirection(direction)
		matrix[x][y] = 'X'
		return simulateGuardPath(matrix, x, y, newDirection, visited)
	}

	matrix[newX][newY] = direction
	matrix[x][y] = 'X'
	return simulateGuardPath(matrix, newX, newY, newDirection, visited)
}

func part2() {
	matrix, startX, startY, direction := readFile()

	total := 0

	for x, rows := range matrix {
		for y, char := range rows {
			if char == '.' {
				newMatrix := make([][]rune, len(matrix))
				for i := range matrix {
					newMatrix[i] = append([]rune(nil), matrix[i]...)
				}

				newMatrix[x][y] = '#'

				_, foundLoop := simulateGuardPath(newMatrix, startX, startY, direction, make(map[string]bool))
				if foundLoop {
					total++
				}
			}
		}
	}

	fmt.Println("Answer to part 2:", total)

}

func main() {
	// part1()
	part2()
}
