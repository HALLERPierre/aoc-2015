package main

import (
	_ "embed"
	"fmt"
	"log"
)

//go:embed input
var input string

func main() {
	floor := puzzle1(input)
	fmt.Printf("Floor is %v\n", floor)

	basement_position := puzzle2(input)
	fmt.Printf("Enter basement at %v\n", basement_position)
}

func puzzle1(input string) int {
	floor := 0
	for _, ch := range input {
		switch ch {
		case '(':
			floor += 1
		case ')':
			floor += -1
		}
	}
	return floor
}

func puzzle2(input string) int {
	floor := 0

	for i, ch := range input {
		switch ch {
		case '(':
			floor += 1
		case ')':
			floor += -1
		}
		if floor < 0 {
			return i + 1
		}
	}

	log.Fatal("Santa never go in basement")
	return 0
}
