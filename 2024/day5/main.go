package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func MultipleLines(fileName string) ([]string, error) {
	inputSlice := make([]string, 0)
	file, err := os.Open(fileName)
	if err != nil {
		return inputSlice, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputSlice = append(inputSlice, scanner.Text())
	}
	err = file.Close()
	if err != nil {
		return inputSlice, err
	}
	return inputSlice, nil
}

func abs(distance int) int {
	if distance < 0 {
		return -distance
	}
	return distance
}

func check(rules map[int][]int, instructions []int) int {
	for index, number := range instructions {
		for _, prevNumber := range instructions[:index] {
			if slices.Contains(rules[number], prevNumber) {
				return 0
			}
		}
	}
	return instructions[len(instructions)/2]
}

func main() {

	lines, err := MultipleLines("input.txt")

	if err != nil {
		fmt.Print("dupa")
	}

	var prev = 0
	var next = 0

	var orderRules map[int][]int
	orderRules = make(map[int][]int)

	var i = 0
	var sum = 0

	for index, line := range lines {
		dupa := strings.Split(line, "|")
		if len(dupa) == 1 {
			i = index
			break
		}

		prev, err = strconv.Atoi(dupa[0])
		next, err = strconv.Atoi(dupa[1])
		orderRules[prev] = append(orderRules[prev], (next))
	}

	for _, line := range lines[i+1:] {
		dupa := strings.Split(line, ",")
		var numbers []int

		for _, number := range dupa {
			var num, _ = strconv.Atoi(number)
			numbers = append(numbers, num)
		}

		var num = check(orderRules, numbers)

		sum += num
	}
	fmt.Println(sum)

}
