package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ColorMap map[string]int

func getGameId(text string) (int, error) {
	gameParts := strings.Split(text, " ")
	gameID, err := strconv.Atoi(gameParts[1])

	if err != nil {
		return 0, err
	}

	return gameID, nil
}

func isColorValid(color string, ballCount int, maxBallsPerColor ColorMap) bool {
	maxAllowed, found := maxBallsPerColor[color]
	return !found || ballCount <= maxAllowed
}

func transformColorCount(colorCount string) (int, string) {
	countAndColor := strings.Split(colorCount, " ")
	ballCount, err := strconv.Atoi(countAndColor[0])
	if err != nil {
		panic("Error converting ball count:")
	}

	color := countAndColor[1]

	return ballCount, color
}

func updateMinimalBallsCount(minBallsPerColor ColorMap, ballCount int, color string) ColorMap {
	currentMinAmount, found := minBallsPerColor[color]

	if found {
		if currentMinAmount < ballCount {
			minBallsPerColor[color] = ballCount
		}
	}

	return minBallsPerColor
}

func calcualatePowerOfSetCubes(minBallsPerColor ColorMap) int {
	power := 1

	for _, v := range minBallsPerColor {
		power *= v
	}

	return power
}

func calculateGameStats(data string) (bool, int) {
	gameRounds := strings.Split(data, "; ")
	maxBallsPerColor := map[string]int{"red": 12, "blue": 13, "green": 13}
	minBallsPerColor := map[string]int{"red": 0, "blue": 0, "green": 0}

	isGameValid := true

	for _, round := range gameRounds {
		colorCounts := strings.Split(round, ", ")

		for _, colorCount := range colorCounts {
			ballCount, color := transformColorCount(colorCount)
			minBallsPerColor = updateMinimalBallsCount(minBallsPerColor, ballCount, color)

			if !isColorValid(color, ballCount, maxBallsPerColor) {
				isGameValid = false
			}
		}
	}

	return isGameValid, calcualatePowerOfSetCubes(minBallsPerColor)
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	defer file.Close()
	totalValid := 0
	totalPower := 0

	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, ":")
		gameInfo := strings.TrimSpace(parts[0])
		gameData := strings.TrimSpace(parts[1])

		gameID, err := getGameId(gameInfo)

		if err != nil {
			fmt.Println("Error parsting game id")
			return
		}

		isValid, power := calculateGameStats(gameData)
		if isValid {
			totalValid += gameID
		}

		totalPower += power

	}

	fmt.Println("total Power ", totalPower)
	fmt.Println("total valid games ", totalValid)
}
