package day2

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readAndReturn(filePath string) [][]string {

	file, err := os.Open("day2/" + filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineArr := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Fields(line)
		if len(splitLine) == 0 {
			continue
		}

		if err != nil {
			fmt.Println(err)
			return nil
		}
		lineArr = append(lineArr, splitLine)
	}
	return lineArr
}

func Day2Part1(filePath string) {

	lineArr := readAndReturn(filePath)
	count := 0
	for _, line := range lineArr {
		windowedLine := windowed(line, 2)
		isDiffOutofBounds := slices.ContainsFunc(windowedLine, func(window []string) bool {
			return ContainsDiffGreaterThan(window, func(current int) bool {
				abs := math.Abs(float64(current))
				return abs < 1 || abs > 3
			})
		})

		isPositive, _ := windowMatches(windowedLine, func(window []string) bool {
			return ContainsDiffGreaterThan(window, func(current int) bool {
				return current > 0
			})
		})
		isNegative, _ := windowMatches(windowedLine, func(window []string) bool {
			return ContainsDiffGreaterThan(window, func(current int) bool {
				return current < 0
			})
		})
		isSafe := !isDiffOutofBounds && (isPositive || isNegative)
		if isSafe {
			count++
		}
	}

	fmt.Printf("The amount of safe reports is: %v\n", count)
}

func Day2Part2(filePath string) {

	lineArr := readAndReturn(filePath)
	count := 0
	for _, line := range lineArr {
		windowedLine := windowed(line, 2)
		isDiffOutofBounds := false
		errPos := slices.IndexFunc(windowedLine, func(window []string) bool {
			return ContainsDiffGreaterThan(window, func(current int) bool {
				abs := math.Abs(float64(current))
				return abs < 1 || abs > 3
			})
		})
		if errPos != -1 {
			lineWithoutError := append(line[:errPos], line[errPos+1:]...)
			windowdLineWithoutError := windowed(lineWithoutError, 2)
			isDiffOutofBounds = slices.ContainsFunc(windowdLineWithoutError, func(window []string) bool {
				return ContainsDiffGreaterThan(window, func(current int) bool {
					abs := math.Abs(float64(current))
					return abs < 1 || abs > 3
				})
			})
		}

		isPositive, errPos := windowMatches(windowedLine, func(window []string) bool {
			return ContainsDiffGreaterThan(window, func(current int) bool {
				return current > 0
			})
		})
		if !isPositive {
			lineWithoutError := append(line[:errPos], line[errPos+1:]...)
			windowdLineWithoutError := windowed(lineWithoutError, 2)
			isPositive, _ = windowMatches(windowdLineWithoutError, func(window []string) bool {
				return ContainsDiffGreaterThan(window, func(current int) bool {
					return current > 0
				})
			})
		}
		isNegative, errPos := windowMatches(windowedLine, func(window []string) bool {
			return ContainsDiffGreaterThan(window, func(current int) bool {
				return current < 0
			})
		})
		if !isNegative {
			lineWithoutError := append(line[:errPos], line[errPos+1:]...)
			windowdLineWithoutError := windowed(lineWithoutError, 2)
			isNegative, _ = windowMatches(windowdLineWithoutError, func(window []string) bool {
				return ContainsDiffGreaterThan(window, func(current int) bool {
					return current < 0
				})
			})
		}
		isSafe := !isDiffOutofBounds && (isPositive || isNegative)
		if isSafe {
			count++
		}
	}
	fmt.Printf("The amount of safe reports is: %v\n", count)
}

func windowed(slice []string, size int) [][]string {
	var result [][]string

	for i := 0; i <= len(slice)-size; i += 1 {
		result = append(result, slice[i:i+size])
	}

	return result
}

func windowMatches(w [][]string, match func([]string) bool) (bool, int) {
	for i, value := range w {
		if !match(value) {
			return false, i
		}
	}
	return true, 0
}

func ContainsDiffGreaterThan(w []string, compare func(current int) bool) bool {
	first, err := strconv.Atoi(w[0])
	second, err := strconv.Atoi(w[1])
	if err != nil {
		fmt.Println(err)
	}
	diff := second - first
	return compare(diff)
}
