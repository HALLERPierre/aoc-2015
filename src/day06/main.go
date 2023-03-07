package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input
var input string

const (
	On     string = "turn on"
	Off    string = "turn off"
	Toggle string = "toggle"
)

func main() {
	lights, luminosity := getLights()
	lightsOn := countLights(lights)
	luminosityCount := countLuminosity(luminosity)
	fmt.Printf("There is %v lights on\n", lightsOn)
	fmt.Printf("There is %v luminosity\n", luminosityCount)
}

func countLights(lights [1000][1000]bool) int {
	count := 0

	for _, line := range lights {
		for _, light := range line {
			if light {
				count += 1
			}
		}
	}
	return count
}

func countLuminosity(lights [1000][1000]int) int {
	count := 0

	for _, line := range lights {
		for _, luminosity := range line {
			count += luminosity
		}
	}
	return count
}

func getLights() ([1000][1000]bool, [1000][1000]int) {
	var lights [1000][1000]bool
	var luminosity [1000][1000]int
	for x, line := range luminosity {
		for y := range line {
			luminosity[x][y] = 0
		}
	}

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		re := regexp.MustCompile(`(?P<x_start>\d+),(?P<y_start>\d+) through (?P<x_end>\d+),(?P<y_end>\d+)`)
		matches := re.FindStringSubmatch(line)

		xStart, _ := strconv.Atoi(matches[re.SubexpIndex("x_start")])
		yStart, _ := strconv.Atoi(matches[re.SubexpIndex("y_start")])
		xEnd, _ := strconv.Atoi(matches[re.SubexpIndex("x_end")])
		yEnd, _ := strconv.Atoi(matches[re.SubexpIndex("y_end")])

		var action string
		if strings.Contains(line, On) {
			action = "on"
		}
		if strings.Contains(line, Off) {
			action = "off"
		}
		if strings.Contains(line, Toggle) {
			action = "toggle"
		}

		for x := xStart; x <= xEnd; x++ {
			for y := yStart; y <= yEnd; y++ {
				switch action {
				case "on":
					lights[x][y] = true
					luminosity[x][y] += 1
				case "off":
					lights[x][y] = false
					if luminosity[x][y] > 0 {
						luminosity[x][y] -= 1
					}
				case "toggle":
					lights[x][y] = !lights[x][y]
					luminosity[x][y] += 2
				}
			}
		}
	}
	return lights, luminosity
}
