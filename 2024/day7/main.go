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

func concatInts(a, b int) int {
	x := strconv.Itoa(a)
	y := strconv.Itoa(b)

	result, err := strconv.Atoi(x + y)
	if err != nil {
		fmt.Println("Failed to concat strings", a, b)
	}

	return result
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
	if node.index == len(node.nums) && node.value == node.expectedResult {
		fmt.Println(node.expectedResult, node.nums)
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

	// This is additional option for part 2
	concatNode := &TreeNode{
		value:          concatInts(node.value, nextNode),
		index:          node.index + 1,
		nums:           node.nums,
		expectedResult: node.expectedResult,
	}

	// Remove concatNode for part1 answer
	return checkPossibleEquations(addNode) || checkPossibleEquations(multiplyNode) || checkPossibleEquations(concatNode)
}

func part2() {
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

	fmt.Println("Answer to part 2:", total)

}

func main() {
	part2()
}
