package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OrderingRules [][]int
type CoordinateList [][]int

func readFile() (OrderingRules, CoordinateList) {
	file, err := os.Open("day5.test.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var orderingRules OrderingRules
	var coordinateList CoordinateList
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			var rule []int
			for _, part := range parts {
				num, _ := strconv.Atoi(part)
				rule = append(rule, num)
			}

			orderingRules = append(orderingRules, rule)
		} else if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			var coordinate []int
			for _, part := range parts {
				num, _ := strconv.Atoi(part)
				coordinate = append(coordinate, num)
			}

			coordinateList = append(coordinateList, coordinate)
		}

	}

	return orderingRules, coordinateList

}

func part1() {
	orderingRules, coordinateList := readFile()

	fmt.Println("Part 1 answer:", orderingRules)
	fmt.Println("Part 2 answer:", coordinateList)
}

func main() {
	part1()
}
