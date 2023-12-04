package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func getFirstLastCombo(word string) int {
	//Find first number and multiply by 10
	var first int
	var last int
	for i := 0; i < len(word); i++ {
		if unicode.IsDigit(rune(word[i])) {
			first = int(word[i]-'0') * 10
			break
		}
	}
	//Find last number and add
	for i := len(word) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(word[i])) {
			last = int(word[i] - '0')
			break
		}
	}
	fmt.Println(word)
	combo := first + last
	fmt.Println(combo)
	return combo
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		sum += getFirstLastCombo(fileScanner.Text())
	}
	fmt.Println(sum)
}
