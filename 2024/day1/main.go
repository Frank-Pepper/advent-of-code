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

func part1(left []int, right []int, size int) int {
	var total = 0
	for i := 0; i < size; i++ {
		distance := left[i] - right[i]
		total += abs(distance)
	}
	return total
}

func part2(left []int, right []int, size int) int {
	var rmap map[int]int
	rmap = make(map[int]int)
	for i := 0; i < size; i++ {
		rmap[right[i]]++
	}
	var result = 0
	for i := 0; i < size; i++ {
		result += rmap[left[i]] * left[i]
	}
	return result
}

func main() {

	lines, err := MultipleLines("input.txt")

	if err != nil {
		fmt.Print("dupa")
	}

	fmt.Println(len(lines))

	size := len(lines)
	left := make([]int, size)
	right := make([]int, size)

	for index, line := range lines {
		dupa := strings.Fields(line)
		left[index], err = strconv.Atoi(dupa[0])
		right[index], err = strconv.Atoi(dupa[1])
	}

	slices.Sort(left)
	slices.Sort(right)

	fmt.Println(part1(left, right, size))

	fmt.Println(part2(left, right, size))
}
