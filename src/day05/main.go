package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/emirpasic/gods/lists/arraylist"
)

//go:embed input
var input string

func main() {
	correctLinesPartOne := 0
	correctLinesPartTwo := 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		if isCorrectPartOne(line) {
			correctLinesPartOne += 1
		}
		if isCorrectPartTwo(line) {
			fmt.Printf("%v \n", line)
			correctLinesPartTwo += 1
		}
	}
	fmt.Printf("Part One: There is %v correct lines\n", correctLinesPartOne)
	fmt.Printf("Part Two: There is %v correct lines\n", correctLinesPartTwo)
}

func isCorrectPartOne(line string) bool {
	vowels := arraylist.New('a', 'e', 'i', 'o', 'u')
	forbiddenStrings := []string{"ab", "cd", "pq", "xy"}
	countVowels := 0
	var prevChar rune
	hasDoubleChar := false

	for _, char := range line {
		if vowels.Contains(char) {
			countVowels += 1
		}
		if prevChar == char {
			hasDoubleChar = true
		}
		prevChar = char
	}

	for _, forbiddenString := range forbiddenStrings {
		if strings.Contains(line, forbiddenString) {
			return false
		}
	}

	return countVowels >= 3 && hasDoubleChar
}

func isCorrectPartTwo(line string) bool {
	repeatChars := false
	repeatCharBetween := true
	var prevChar rune
	for i, char := range line {
		if i != 0 {
			var sb strings.Builder
			sb.WriteRune(prevChar)
			sb.WriteRune(char)
			rightString := line[i+1:]
			if strings.Contains(rightString, sb.String()) {
				repeatChars = true
			}
		}
		if i+2 >= len(line) {
			continue
		}
		oneLetterBetween := []rune(line)[i+2]
		if char == oneLetterBetween {
			repeatCharBetween = true
		}
		prevChar = char
	}

	return repeatChars && repeatCharBetween
}
