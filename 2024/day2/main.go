package main

import (
	"bufio"
	"fmt"
	"os"
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

func intArray(numbers []string) []int {
	ints := make([]int, 0)

	for _, number := range numbers {
		num, _ := strconv.Atoi(number)
		ints = append(ints, num)
	}
	return ints
}

func abs(distance int) int {
	if distance < 0 {
		return -distance
	}
	return distance
}

func isSafe(array []int) int {
	var size = len(array)
	var total = 0
	var prevDistance = 0
	for i := 0; i < size-1; i++ {
		distance := array[i] - array[i+1]
		total += abs(distance)
		if prevDistance*distance < 0 {
			return 0
		}
		if 1 > abs(distance) || abs(distance) > 3 {
			return 0
		}
		prevDistance = distance
	}
	return 1
}

func part1(lines []string) int {
	var totalSafe = 0
	for _, line := range lines {
		dupa := intArray(strings.Fields(line))
		totalSafe += isSafe(dupa)
	}
	return totalSafe
}

func main() {

	lines, _ := MultipleLines("input.txt")

	fmt.Println(len(lines))

	println(part1(lines))

}
