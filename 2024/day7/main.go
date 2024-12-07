package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	result int
	nums   []int
}

func convertStringToIntArray(parts string) []int {
	var nums []int

	for _, num := range strings.Fields(parts) {
		convertedNum, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Smth went wrong", err)
			fmt.Println(err)
		}
		nums = append(nums, convertedNum)
	}

	return nums
}

func readFile() []Equation {
	file, err := os.Open("day7.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	defer file.Close()
	var equations []Equation

	for scanner.Scan() {
		parts := scanner.Text()
		splittedParts := strings.Split(parts, ":")
		result, err := strconv.Atoi(splittedParts[0])
		if err != nil {
			fmt.Println("Smth went wrong", err)
			panic("Failed to convert first number")
		}

		equations = append(equations, Equation{
			result: result,
			nums:   convertStringToIntArray(splittedParts[1]),
		})

	}

	return equations

}

type TreeNode struct {
	value          int
	index          int
	nums           []int
	expectedResult int
}

func checkPossibleEquations(node *TreeNode) bool {
	if node.value == node.expectedResult {
		return true
	}

	if node.index >= len(node.nums) {
		return false
	}

	nextNode := node.nums[node.index]

	addNode := &TreeNode{
		value:          node.value + nextNode,
		index:          node.index + 1,
		nums:           node.nums,
		expectedResult: node.expectedResult,
	}

	multiplyNode := &TreeNode{
		value:          node.value * nextNode,
		index:          node.index + 1,
		nums:           node.nums,
		expectedResult: node.expectedResult,
	}

	return checkPossibleEquations(addNode) || checkPossibleEquations(multiplyNode)
}

func part1() {
	equations := readFile()

	total := 0
	for _, equation := range equations {

		treeNode := &TreeNode{
			value:          equation.nums[0],
			index:          1,
			nums:           equation.nums,
			expectedResult: equation.result,
		}

		if checkPossibleEquations(treeNode) {
			total += equation.result
		}
	}

	fmt.Println("Answer to part 1:", total)

}

func main() {
	part1()
}
