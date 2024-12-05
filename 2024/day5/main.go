package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type IntGrid [][]int
type OrderingRules map[int][]int

func readFile() (IntGrid, IntGrid) {
	file, err := os.Open("day5.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var orderingRules IntGrid
	var updates IntGrid
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

			updates = append(updates, coordinate)
		}

	}

	return orderingRules, updates

}

func mapRules(rules IntGrid) OrderingRules {
	var rulesMapping = make(OrderingRules)
	for _, rule := range rules {
		y, x := rule[0], rule[1]
		rulesMapping[x] = append(rulesMapping[x], y)
	}

	return rulesMapping
}

func isUpdateValid(update []int, rules OrderingRules) bool {
	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			if slices.Contains(rules[update[i]], update[j]) {
				return false

			}
		}
	}
	return true

}

func reorderUpdates(update []int, rules OrderingRules) []int {
	reordered := make([]int, len(update))
	copy(reordered, update)

	for i := 0; i < len(reordered); i++ {
		for j := i + 1; j < len(reordered); j++ {
			if slices.Contains(rules[reordered[i]], reordered[j]) {
				reordered[i], reordered[j] = reordered[j], reordered[i]
			}
		}
	}

	return reordered
}

func part1() {
	orderingRulesData, updates := readFile()
	orderingRules := mapRules(orderingRulesData)

	total := 0

	for _, update := range updates {
		if isUpdateValid(update, orderingRules) {
			total += update[len(update)/2]
		}
	}

	fmt.Println("Answer to part 1", total)
}

func part2() {
	orderingRulesData, updates := readFile()
	orderingRules := mapRules(orderingRulesData)

	total := 0

	for _, update := range updates {
		if !isUpdateValid(update, orderingRules) {
			reorderedUpdate := reorderUpdates(update, orderingRules)
			total += reorderedUpdate[len(reorderedUpdate)/2]
		}
	}

	fmt.Println("Answer to part 2", total)
}

func main() {
	part1()
	part2()
}
