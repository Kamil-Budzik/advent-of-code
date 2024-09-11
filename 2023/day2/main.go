package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getGameId(text string) (int, error) {
	gameParts := strings.Split(text, " ")
	gameID, err := strconv.Atoi(gameParts[1])

	if err != nil {
		return 0, err
	}

	return gameID, nil
}

func isColorValid(color string, ballCount int, maxBallsPerColor map[string]int) bool {
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

func checkValidGame(data string) bool {
	gameRounds := strings.Split(data, "; ")
	maxBallsPerColor := map[string]int{"red": 12, "blue": 13, "green": 13}
	isGameValid := true

	for _, round := range gameRounds {
		colorCounts := strings.Split(round, ", ")

		for _, colorCount := range colorCounts {
			ballCount, color := transformColorCount(colorCount)

			if !isColorValid(color, ballCount, maxBallsPerColor) {
				isGameValid = false
			}
		}
	}
	return isGameValid
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	defer file.Close()
	total := 0

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

		isValid := checkValidGame(gameData)
		if isValid {
			total += gameID
		}

	}

	fmt.Print(total)
}
