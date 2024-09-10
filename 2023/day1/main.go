package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	defer file.Close()
	total := 0

	for scanner.Scan() {
		var firstDigit, lastDigit rune
		foundFirst := false
		text := scanner.Text()

		for _, r := range text {
			if unicode.IsDigit(r) {
				if !foundFirst {
					firstDigit = r
					foundFirst = true
				}
				lastDigit = r
			}
		}

		first := int(firstDigit - '0')
		last := int(lastDigit - '0')
		number := first*10 + last
		total += number
	}

	fmt.Println(total)

}
