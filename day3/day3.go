package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readAndReturn(filePath string) []string {

	file, err := os.Open("day3/" + filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func Day3Part1(filePath string) {
	lines := readAndReturn(filePath)
	answer := 0
	for _, line := range lines {
		mulsRegex := regexp.MustCompile(`(mul)(\(\d+,\d+\))`)
		muls := mulsRegex.FindAll([]byte(line), -1)
		for _, mulByte := range muls {
			numberRegex := regexp.MustCompile(`\d+`)
			numbers := numberRegex.FindAll(mulByte, -1)
			first, err := strconv.Atoi(string(numbers[0]))
			second, err := strconv.Atoi(string(numbers[1]))
			if err != nil {
				fmt.Println(err)
				return
			}
			answer += first * second
		}
	}
	fmt.Printf("The sum of mul instructions is: %v\n", answer)
}

func Day3Part2(filePath string) {
	lines := readAndReturn(filePath)
	answer := 0
	isEnabled := true
	for _, line := range lines {
		doesRegex := regexp.MustCompile(`(do\(\)|don't\(\))`)
		doesIndices := doesRegex.FindAllIndex([]byte(line), -1)
		startPosition := 0
		remainingLine := strings.Clone(line)
		for _, indexes := range doesIndices {

			if isEnabled {
				lineToCheck := line[startPosition:indexes[0]]
				multiplication := sumAllMulsInString(lineToCheck)
				if multiplication != -1 {
					answer += multiplication
				}
			}
			startPosition = indexes[1]

			match := line[indexes[0]:indexes[1]]
			isDo := match == "do()"
			remainingLine = line[indexes[1]:]
			if !isDo {
				isEnabled = false
			} else {
				isEnabled = true
			}
		}
		if isEnabled {
			answer += sumAllMulsInString(remainingLine)
		}
	}
	fmt.Printf("The sum of mul instructions is: %v\n", answer)
}

func sumAllMulsInString(str string) int {
	mulsRegex := regexp.MustCompile(`(mul)(\(\d+,\d+\))`)
	muls := mulsRegex.FindAll([]byte(str), -1)
	sum := 0
	for _, mulByte := range muls {
		numberRegex := regexp.MustCompile(`\d+`)
		numbers := numberRegex.FindAll(mulByte, -1)
		first, err := strconv.Atoi(string(numbers[0]))
		second, err := strconv.Atoi(string(numbers[1]))
		if err != nil {
			fmt.Println(err)
			return -1
		}
		sum += first * second
	}
	return sum
}
