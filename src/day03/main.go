package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/emirpasic/gods/sets/hashset"
)

const (
	Up    string = "^"
	Down         = "v"
	Left         = "<"
	Right        = ">"
)

type Coord struct {
	x int
	y int
}

//go:embed input
var input string

func main() {
	directions := parseInput()
	visitedOnce := puzzle1(directions)
	fmt.Printf("Visited once: %v\n", visitedOnce)
	visitedOnceWithRobo := puzzle2(directions)
	fmt.Printf("Visited once with robo: %v\n", visitedOnceWithRobo)
}

func puzzle1(directions []string) int {
	currentCoord := Coord{0, 0}
	visited := hashset.New(Coord{0, 0})
	for _, direction := range directions {
		switch direction {
		case Up:
			currentCoord = Coord{currentCoord.x, currentCoord.y + 1}
		case Down:
			currentCoord = Coord{currentCoord.x, currentCoord.y - 1}
		case Left:
			currentCoord = Coord{currentCoord.x - 1, currentCoord.y}
		case Right:
			currentCoord = Coord{currentCoord.x + 1, currentCoord.y}
		}
		visited.Add(currentCoord)
	}
	return visited.Size()
}

func puzzle2(directions []string) int {
	santaCoord := Coord{0, 0}
	roboCoord := Coord{0, 0}
	visited := hashset.New(Coord{0, 0})
	for index, direction := range directions {
		currentCoord := santaCoord
		if index%2 == 1 {
			currentCoord = roboCoord
		}
		switch direction {
		case Up:
			currentCoord = Coord{currentCoord.x, currentCoord.y + 1}
		case Down:
			currentCoord = Coord{currentCoord.x, currentCoord.y - 1}
		case Left:
			currentCoord = Coord{currentCoord.x - 1, currentCoord.y}
		case Right:
			currentCoord = Coord{currentCoord.x + 1, currentCoord.y}
		}
		visited.Add(currentCoord)

		if index%2 == 0 {
			santaCoord = currentCoord
		} else {
			roboCoord = currentCoord
		}
	}
	return visited.Size()
}

func parseInput() []string {
	directions := []string{}

	for _, char := range strings.Split(input, "") {
		var newChar string
		switch char {
		case Up:
			newChar = Up
		case Down:
			newChar = Down
		case Left:
			newChar = Left
		case Right:
			newChar = Right
		}
		if newChar != "" {
			directions = append(directions, newChar)
		}
	}
	return directions
}
