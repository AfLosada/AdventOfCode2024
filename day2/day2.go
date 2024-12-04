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
		isDiffInbounds := true
		outOfBoundsErrIndex := slices.IndexFunc(windowedLine, func(window []string) bool {
			return ContainsDiffGreaterThan(window, func(current int) bool {
				abs := math.Abs(float64(current))
				return abs < 1 || abs > 3
			})
		})
		if outOfBoundsErrIndex != -1 {
			fmt.Printf("Removed for operation: %v\n index: %v\n value: %v\n", "outOfBouds", outOfBoundsErrIndex, line[outOfBoundsErrIndex])
			lineWithoutError := sliceWithErrorIndex(line, outOfBoundsErrIndex)
			lineWithoutError2 := sliceWithErrorIndex(line, outOfBoundsErrIndex+1)
			windowedLineWithoutError := windowed(lineWithoutError, 2)
			windowedLineWithoutError2 := windowed(lineWithoutError2, 2)

			isDiffOutofBounds1 := slices.ContainsFunc(windowedLineWithoutError, func(window []string) bool {
				return ContainsDiffGreaterThan(window, func(current int) bool {
					abs := math.Abs(float64(current))
					return abs < 1 || abs > 3
				})
			})
			isDiffOutofBounds2 := slices.ContainsFunc(windowedLineWithoutError2, func(window []string) bool {
				return ContainsDiffGreaterThan(window, func(current int) bool {
					abs := math.Abs(float64(current))
					return abs < 1 || abs > 3
				})
			})
			isDiffInbounds = !isDiffOutofBounds1 || !isDiffOutofBounds2
			if isDiffInbounds {
				windowedLine = windowedLineWithoutError
			}
		}

		isPositive, _ := runCallbackFailure(line, windowedLine, func(w [][]string) (bool, int) {
			return windowMatches(w, func(window []string) bool {
				return ContainsDiffGreaterThan(window, func(current int) bool {
					return current > 0
				})
			})
		}, "positive")

		isNegative, _ := runCallbackFailure(line, windowedLine, func(w [][]string) (bool, int) {
			return windowMatches(w, func(window []string) bool {
				return ContainsDiffGreaterThan(window, func(current int) bool {
					return current < 0
				})
			})
		}, "negative")

		isSafe := isDiffInbounds && (isPositive || isNegative)
		if isSafe {
			fmt.Println(line)
			fmt.Printf(" INBOUNDS: %v\n POSITIVE: %v\n NEGATIVE: %v\n", isDiffInbounds, isPositive, isNegative)
			count++
		}
		fmt.Println("=======================================")
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

func runCallbackFailure(line []string, w [][]string, callback func(w [][]string) (bool, int), name string) (bool, int) {
	pass, errIndex := callback(w)
	if !pass {
		newLine := sliceWithErrorIndex(line, errIndex)
		newLine2 := sliceWithErrorIndex(line, errIndex+1)
		windowedLine := windowed(newLine, 2)
		windowedLine2 := windowed(newLine2, 2)
		answ1, i1 := callback(windowedLine)
		answ2, i2 := callback(windowedLine2)
		fmt.Printf("Removed for operation: %v\n index: %v\n value: %v\n", name, errIndex, line[errIndex])
		if answ1 {
			return answ1, i1
		}
		if answ2 {
			return answ2, i2
		}
	}
	return pass, errIndex
}

func sliceWithErrorIndex(line []string, errIndex int) []string {
	leftPart := line[:errIndex]
	leftPartCopy := make([]string, len(leftPart))
	copy(leftPartCopy, leftPart)
	rightPart := line[errIndex+1:]
	rightPartCopy := make([]string, len(rightPart))
	copy(rightPartCopy, rightPart)
	return append(leftPartCopy, rightPartCopy...)
}
