package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func getFirstLastCombo(word string, numberMap map[string]int) int {
	//Find first number and multiply by 10
	var first int
	var last int
	var numberWord string
	fmt.Println("\n\n" + word)
	for i := 0; i < len(word); i++ {
		fmt.Println(numberWord)
		numberWord += string(word[i])
		if unicode.IsDigit(rune(word[i])) {
			first = int(word[i]-'0') * 10
			break
		} else if substring, found := findSubstringInMap(word, numberMap); found {
			first = numberMap[substring] * 10
			break
		}
	}
	//Find last number and add
	numberWord = ""
	for i := len(word) - 1; i >= 0; i-- {
		numberWord += string(word[i])
		if unicode.IsDigit(rune(word[i])) {
			last = int(word[i] - '0')
			break
		} else if substring, found := findSubstringInMap(reverseString(numberWord), numberMap); found {
			last = numberMap[substring]
			break
		}
	}
	combo := first + last
	fmt.Println(combo)
	return combo
}

func reverseString(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	print(string(runes))
	print("\n")
	return string(runes)
}

func findSubstringInMap(word string, numberMap map[string]int) (string, bool) {
	for substring := range numberMap {
		if strings.Contains(word, substring) {
			return substring, true
		}
	}
	return "", false
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	numberMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	sum := 0
	for fileScanner.Scan() {
		sum += getFirstLastCombo(fileScanner.Text(), numberMap)
	}
	fmt.Println(sum)
}
