package main

import (
	"aoc-2015/utils"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input
var input string

type Dimension struct {
	l int
	w int
	h int
}

func main() {
	dimensions := parseInput()
	sqf := puzzle1(dimensions)
	fmt.Printf("square feet: %v\n", sqf)
	ribbon := puzzle2(dimensions)
	fmt.Printf("ribbon feet: %v\n", ribbon)
}

func parseInput() []Dimension {
	dimensions := []Dimension{}

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		re := regexp.MustCompile(`(?P<l>\d+)x(?P<w>\d+)x(?P<h>\d+)`)
		matches := re.FindStringSubmatch(line)
		l, _ := strconv.Atoi(matches[re.SubexpIndex("l")])
		w, _ := strconv.Atoi(matches[re.SubexpIndex("w")])
		h, _ := strconv.Atoi(matches[re.SubexpIndex("h")])

		dimensions = append(dimensions, Dimension{l, w, h})
	}
	return dimensions
}

func puzzle1(dimensions []Dimension) int {
	sqf := 0
	for _, d := range dimensions {
		sqf += 2*d.l*d.w + 2*d.w*d.h + 2*d.h*d.l
		sqf += utils.Min(d.l*d.w, d.w*d.h, d.h*d.l)
	}
	return sqf
}

func puzzle2(dimensions []Dimension) int {
	ribbonFeet := 0
	for _, d := range dimensions {
		ribbonFeet += utils.Min((d.l+d.w)*2, (d.l+d.h)*2, (d.w+d.h)*2)
		ribbonFeet += d.l * d.w * d.h
	}
	return ribbonFeet
}
