package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/maps/hashmap"
)

//go:embed input
var input string

type Instruction struct {
	sources []string
	action  string
	value   int
}

func main() {
	instructions := parseInstructions()
	values := hashmap.New()
	value := findValue(instructions, values, "a")
	fmt.Printf("Value of a wire: %v\n", value)

	instructions.Put("b", Instruction{[]string{}, "SET", value})
	newValues := hashmap.New()
	newValue := findValue(instructions, newValues, "a")
	fmt.Printf("New value of a wire: %v\n", newValue)
}

func findValue(instructions *hashmap.Map, values *hashmap.Map, target string) int {
	instructionI, _ := instructions.Get(target)
	instruction := instructionI.(Instruction)
	action := instruction.action

	value, exist := values.Get(target)
	if exist {
		return value.(int)
	}
	var newValue int
	switch action {
	case "AND":
		if len(instruction.sources) == 1 {
			newValue = instruction.value & findValue(instructions, values, instruction.sources[0])
		} else {
			newValue = findValue(instructions, values, instruction.sources[0]) & findValue(instructions, values, instruction.sources[1])
		}
		break
	case "OR":
		if len(instruction.sources) == 1 {
			newValue = instruction.value & findValue(instructions, values, instruction.sources[0])
		} else {
			newValue = findValue(instructions, values, instruction.sources[0]) | findValue(instructions, values, instruction.sources[1])
		}
		break
	case "NOT":
		newValue = 65535 ^ findValue(instructions, values, instruction.sources[0])
		break
	case "LSHIFT":
		newValue = findValue(instructions, values, instruction.sources[0]) << instruction.value
		break
	case "RSHIFT":
		newValue = findValue(instructions, values, instruction.sources[0]) >> instruction.value
		break
	case "WIRE":
		newValue = findValue(instructions, values, instruction.sources[0])
		break
	case "SET":
		newValue = instruction.value
		break
	default:
		fmt.Printf("Unknown action %v\n", action)
	}
	values.Put(target, newValue)
	return newValue
}

func parseInstructions() *hashmap.Map {
	instructions := hashmap.New()

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		re := regexp.MustCompile(`-> (?P<destination>.*)`)
		destMatches := re.FindStringSubmatch(line)
		destination := destMatches[re.SubexpIndex("destination")]

		andOrRe := regexp.MustCompile(`(?P<node1>.*) (?P<andOr>(AND|OR)) (?P<node2>.*) ->`)
		andOrMatches := andOrRe.FindStringSubmatch(line)
		if len(andOrMatches) > 0 {
			node1 := andOrMatches[andOrRe.SubexpIndex("node1")]
			value, isValue := strconv.Atoi(node1)
			node2 := andOrMatches[andOrRe.SubexpIndex("node2")]
			andOr := andOrMatches[andOrRe.SubexpIndex("andOr")]
			if isValue == nil {
				instructions.Put(destination, Instruction{[]string{node2}, andOr, value})
			} else {
				instructions.Put(destination, Instruction{[]string{node1, node2}, andOr, 0})
			}
			continue
		}

		shiftRe := regexp.MustCompile(`(?P<node1>.*) (?P<shift>(RSHIFT|LSHIFT)) (?P<value>.*) ->`)
		shiftMatches := shiftRe.FindStringSubmatch(line)
		if len(shiftMatches) > 0 {
			node1 := shiftMatches[shiftRe.SubexpIndex("node1")]
			value, _ := strconv.Atoi(shiftMatches[shiftRe.SubexpIndex("value")])
			shift := shiftMatches[shiftRe.SubexpIndex("shift")]
			instructions.Put(destination, Instruction{[]string{node1}, shift, value})
			continue
		}

		notRe := regexp.MustCompile(`NOT (?P<node1>.*) ->`)
		notMatches := notRe.FindStringSubmatch(line)
		if len(notMatches) > 0 {
			node1 := notMatches[notRe.SubexpIndex("node1")]
			instructions.Put(destination, Instruction{[]string{node1}, "NOT", 0})
			continue
		}

		signalRe := regexp.MustCompile(`(?P<signal>\d+) ->`)
		signalMatches := signalRe.FindStringSubmatch(line)
		if len(signalMatches) > 0 {
			signal, _ := strconv.Atoi(signalMatches[signalRe.SubexpIndex("signal")])
			instructions.Put(destination, Instruction{[]string{}, "SET", signal})
			continue
		}

		// it's a wire
		wireRe := regexp.MustCompile(`(?P<wire>.*) ->`)
		wireMatches := wireRe.FindStringSubmatch(line)
		if len(wireMatches) > 0 {
			wire := wireMatches[wireRe.SubexpIndex("wire")]
			instructions.Put(destination, Instruction{[]string{wire}, "WIRE", 0})
			continue
		}
	}

	return instructions
}
